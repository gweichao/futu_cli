package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	libf "libfunc"
	"pbgo/InitConnect"
	"pbgo/KeepAlive"

	"github.com/gogo/protobuf/proto"
	// "github.com/golang/protobuf/proto"
)

var (
	CheckTs libf.CheckTsStu

	debugflag_conn      bool = false
	debugflag_respmng   bool = false
	debugflag_keepalive bool = false
)

/*e.g. see main.go */

// RequestStu is a service for connect futuOpenD
type RequestStu struct {
	ConnAddr string

	FutuID uint64
	ConnID uint64

	Conn net.Conn

	readyFlag uint32

	sync.RWMutex
}

func (this *RequestStu) ReadyGet() bool {

	ready := atomic.LoadUint32(&this.readyFlag)

	if ready != 0 {
		return true
	}

	return false
}

// toSec 超时秒数，默认10秒
func (this *RequestStu) ReadyWait(toSec int) {

	if toSec <= 0 {
		toSec = 10
	}

	wait1 := int64(toSec) * NanoTi64 // ns 1000000000
	nano1 := time.Now().UnixNano()

	for true {
		time.Sleep(time.Millisecond * 200)

		if time.Now().UnixNano()-nano1 >= wait1 {
			break
		}

		if this.ReadyGet() {
			return
		}

	}

	return
}

// Connect futuOpenD
func (this *RequestStu) Connect() (rttnano int64, err error) {

	rttnano = time.Now().UnixNano()
	defer func() {
		rttnano = time.Now().UnixNano() - rttnano
	}()

	atomic.StoreUint32(&this.readyFlag, 0)

	this.Conn, err = net.Dial("tcp", this.ConnAddr)
	if err != nil {
		err = errors.New("cannot connect server, " + err.Error())
		return
	}

	return
}

func (this *RequestStu) sendInitConnect(enterFlagIn int) (rttnano int64, err error) {
	var (
		ProtoID = uint32(1001) // 1001
		pbData  []byte
		n       int
	)
	rttnano = time.Now().UnixNano()
	defer func() {
		rttnano = time.Now().UnixNano() - rttnano
		if err == nil {
			atomic.StoreUint32(&this.readyFlag, 1)
		}
	}()
	if debugflag_conn {
		PrnLog.Debugf("[%v]enter @ %v", enterFlagIn, time.Now().Format("15:04:05.000000"))
	}
	// pack
	pack := &FutuPackStu{}
	pack.ProtoIDSet(ProtoID)
	futuReq := &InitConnect.Request{
		C2S: &InitConnect.C2S{
			ClientID:  ClientId,
			ClientVer: int32(ClientVer),
		},
	}
	pbData, err = proto.Marshal(futuReq)
	if err != nil {
		return
	}
	pack.BodySet(pbData)

	var packResp *FutuPackStu
	packResp, _, n, err = this.Send(pack, true)
	if err != nil {
		PrnLog.Errorf("[%v]Send error:%v", enterFlagIn, err)
		return
	}

	if debugflag_conn {
		PrnLog.Debugf("[%v]n=%v, err=%v @ %v", enterFlagIn, n, err, time.Now().Format("15:04:05.000000"))
	}

	if packResp.ProtoID != ProtoID {
		PrnLog.Warningf("[%v]packResp.ProtoID=%v, send ProtoID=%v", enterFlagIn, packResp.ProtoID, ProtoID)
		return
	}

	if debugflag_conn {
		PrnLog.Debugf("[%v]%v==> @ %v", enterFlagIn, packResp.ProtoID, time.Now().Format("15:04:05.000000"))
	}

	fut := &InitConnect.Response{}
	err = proto.Unmarshal(packResp.body, fut)
	if err != nil {
		PrnLog.Errorf("[%v]unmarshaling error:%v", enterFlagIn, err)
	}

	this.FutuID = fut.S2C.LoginUserID
	this.ConnID = fut.S2C.ConnID

	if fut.GetRetType() != 0 {
		PrnLog.Errorf("[%v]@ %v, Ret=%v, Msg=%v, resp=%v ", enterFlagIn,
			time.Now().Format("15:04:05.000000"),
			fut.GetRetType(), fut.GetRetMsg(), fut.String(),
		)
		return
	}

	PrnLog.Debugf("[%v]success @ %v", enterFlagIn, time.Now().Format("15:04:05.000000") /*, fut.String()*/)

	return
}

func (this *RequestStu) sendkeepAlive(enterFlagIn int, printDetail bool) (rttnano int64, err error) {
	var (
		ProtoID = uint32(1004) // 1004
		pbData  []byte
	)
	rttnano = time.Now().UnixNano()
	defer func() {
		rttnano = time.Now().UnixNano() - rttnano
	}()

	// pack
	pack := &FutuPackStu{}
	pack.ProtoIDSet(ProtoID)
	futuReq := &KeepAlive.Request{
		C2S: &KeepAlive.C2S{
			Time: time.Now().Unix(),
		},
	}
	pbData, err = proto.Marshal(futuReq)
	if err != nil {
		return
	}
	pack.BodySet(pbData)

	var packResp *FutuPackStu
	packResp, _, _, err = this.Send(pack, true)
	if err != nil {
		PrnLog.Errorf("[%v]Send error:%v", enterFlagIn, err)
		return
	}

	if packResp.ProtoID != ProtoID {
		PrnLog.Warningf("[%v]packResp.ProtoID=%v, send ProtoID=%v", enterFlagIn, packResp.ProtoID, ProtoID)
		return
	}

	if debugflag_conn || debugflag_keepalive {
		PrnLog.Debugf("[%v]%v==> @ %v", enterFlagIn, packResp.ProtoID, time.Now().Format("15:04:05.000000"))
	}

	fut := &KeepAlive.Response{}
	err = proto.Unmarshal(packResp.body, fut)
	if err != nil {
		PrnLog.Errorf("[%v]unmarshaling error:%v", enterFlagIn, err)
	}
	if fut.GetRetType() != 0 {
		PrnLog.Errorf("[%v]@ %v, Ret=%v, Msg=%v ", enterFlagIn, time.Now().Format("15:04:05.000000"),
			fut.GetRetType(), fut.GetRetMsg(),
		)
		return
	}
	if printDetail || debugflag_keepalive {
		PrnLog.Debugf("[%v]success req ts=%v, resp ts=%v @ %v", enterFlagIn,
			futuReq.C2S.Time, fut.S2C.Time, time.Now().Format("15:04:05.000000"))
	}

	return
}

/* 1
&this.HeaderFlag   // PackBuild()  -- step3: in RequestStu.Send()
&this.ProtoID      // ProtoIDSet() -- step1
&this.ProtoFmtType // PackBuild()
&this.ProtoVer     // PackBuild()
&this.SerialNo     // PackBuild()
&this.BodyLen      // BodySet()	   -- step2: proto.Marshal(), get *.Request{} []byte, -->pack.BodySet()
&this.BodySHA1     // BodySet()
&this.Reserved
*/
// Send data, Sync 是否同步方式读取响应
func (this *RequestStu) Send(pack *FutuPackStu, SyncIn ...bool) (packResp *FutuPackStu,
	rttnano int64, n int, err error) {
	var (
		packData []byte
		sync1    bool = false
	)
	rttnano = time.Now().UnixNano()
	defer func() {
		rttnano = time.Now().UnixNano() - rttnano
	}()
	if len(SyncIn) > 0 {
		sync1 = SyncIn[0]
	}

	if this.Conn == nil {
		err = errors.New("not init")
		return
	}

	if debugflag_conn {
		PrnLog.Debugf("enter @ %v", time.Now().Format("15:04:05.000000"))
	}

	// pack
	packData, err = pack.PackBuild()
	if err != nil {
		return
	}

	if debugflag_conn {
		PrnLog.Debugf("PackBuild %v @ %v", len(packData), time.Now().Format("15:04:05.000000"))
	}

	// send
	RespPackMng.ChanAdd(pack.SerialNo)
	n, err = this.Conn.Write(packData)
	if err != nil {
		RespPackMng.ChanDel(pack.SerialNo)
		return
	}

	if debugflag_conn {
		PrnLog.Debugf("Conn.Write @ %v", time.Now().Format("15:04:05.000000"))
	}

	if sync1 {
		packResp = RecvSync(pack.SerialNo, 10)
		if packResp == nil {
			err = errors.New("no resp")
		}
	}
	if debugflag_conn {
		PrnLog.Debugf("Recv @ %v", time.Now().Format("15:04:05.000000"))
	}

	return
}

// Recv data
func (this *RequestStu) recvRaw() (pack *FutuPackStu, err error) {

	if this.Conn == nil {
		err = errors.New("not init")
		return
	}

	if debugflag_conn {
		PrnLog.Debugf("enter @ %v", time.Now().Format("15:04:05.000000"))
	}

	// scanner
	scanner := bufio.NewScanner(this.Conn)
	scanner.Buffer([]byte{}, bufio.MaxScanTokenSize*512)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err2 error) {
		if !atEOF && data[0] == 'F' {
			if len(data) > 44 {
				length := uint32(0)
				binary.Read(bytes.NewReader(data[12:16]), binary.LittleEndian, &length)
				if int(length)+4 < len(data) {
					return int(length) + 44, data[:int(length)+44], nil
				}
			}
		}
		return
	})

	if debugflag_conn {
		PrnLog.Debugf("Scan enter @ %v", time.Now().Format("15:04:05.000000"))
	}

	// for scanner.Scan() { // 如果要改成同步调用，每次返回，这里要把 for  改成 if，否则会死循环收包，30秒后超时
	if scanner.Scan() {
		pack = new(FutuPackStu)
		err = pack.PackRead(scanner.Bytes())
		if err != nil {
			err = fmt.Errorf("unpack error, %v", err)
			return
		}
		if debugflag_conn && CheckTs.Check("scanner.Scan.1", 2) {
			PrnLog.Debugf("Scan @ %v", time.Now().Format("15:04:05.000000"))
		}
	}

	if debugflag_conn {
		PrnLog.Debugf("Scan end @ %v", time.Now().Format("15:04:05.000000"))
	}

	if err = scanner.Err(); err != nil {
		err = fmt.Errorf("scann error, %v", err)
	}

	if debugflag_conn {
		PrnLog.Debugf("Scan Err @ %v", time.Now().Format("15:04:05.000000"))
	}

	return
}

// Close connect
func (this *RequestStu) Close() {

	defer atomic.StoreUint32(&this.readyFlag, 0)

	if this.Conn != nil {
		this.Conn.Close()
		this.Conn = nil
	}

}

type ConnManageStu struct {
	request *RequestStu

	connFlag        uint32 // 0 idle, 1 doing
	respPackMngFlag uint32

	sync.RWMutex
}

// Enter() routineConnAlive：Connect() and keep alive 自维护连接和心跳; routineRespPack 收包到中间table；
// debugflag 第一个 打印详情，第二个 跳过实际连接
func (this *ConnManageStu) Enter(RequestIn *RequestStu, debugflagIn ...bool) {

	var (
		enterFlag   int  = libf.GetRand(10000)
		printDetail bool = false
		skipConn    bool = false
	)

	if len(debugflagIn) >= 1 {
		printDetail = debugflagIn[0]
	}
	if len(debugflagIn) >= 2 {
		skipConn = debugflagIn[1]
	}

	if RequestIn == nil {
		PrnLog.Errorf("[%v]no RequestIn @ %v",
			enterFlag, time.Now().Format(TIMEFORMAT_HHMMssMS))
		return
	}

	if this.request != nil {
		this.request.Close()
		this.request = nil
	}
	this.request = RequestIn

	this.routineConnAlive(enterFlag, printDetail, skipConn)
	this.routineRespPack(enterFlag, printDetail)

	return
}

func (this *ConnManageStu) routineRespPack(enterFlagIn int, printDetailIn bool) {

	go func() {

		// 防止重入
		runResponsePackMng1 := atomic.LoadUint32(&this.respPackMngFlag)
		if runResponsePackMng1 != 0 {
			return
		}
		atomic.StoreUint32(&this.respPackMngFlag, 1)

		// init
		RespPackMng.Init()

		PrnLog.Debugf("[%v]enter recvLoop", enterFlagIn)
		ts1 := time.Now()

		// ================ recv and save to response table ================ /
		for {
			if this.request == nil || this.request.Conn == nil {
				time.Sleep(time.Second * 1)
				continue
			}
			packResp, err := this.request.recvRaw()

			if debugflag_respmng {
				PrnLog.Debugf("[%v]recvLoop, ProtoID=%v, SerialNo=%v, err=%v, %v", enterFlagIn,
					packResp.ProtoID, packResp.SerialNo, err, time.Now().Format("15:04:05.000000"))
			}

			if err != nil {
				if CheckTs.Check("recvLoop.1", 120) {
					PrnLog.Errorf("[%v]recvLoop err=%v", enterFlagIn, err)
				}
				continue
			}
			RespPackMng.PackAdd(packResp)
			if time.Since(ts1) > time.Second*300 {
				ts1 = time.Now()
				if debugflag_respmng {
					PrnLog.Debugf("[%v]responsePackMng.Scan, %v", enterFlagIn,
						time.Now().Format("15:04:05.000000"))
				}
				RespPackMng.Scan(0)
			}
		}
		// return
	}()

	return
}

func (this *ConnManageStu) routineConnAlive(enterFlagIn int, printDetail bool, skipConn bool) {

	go func(enterFlag1 int) {
		var (
			connWaitSec  int64 = 5
			KeepAliveSec int64 = int64(SysConfig.KeepAlive)
			errNum       int
		)
		if KeepAliveSec <= 0 {
			KeepAliveSec = 5
		}

		// 防止重入
		connFlag1 := atomic.LoadUint32(&this.connFlag)
		if connFlag1 != 0 {
			return
		}
		atomic.StoreUint32(&this.connFlag, 1) // doing

		PrnLog.Debugf("[%v]enter conn manage", enterFlag1)

		defer func() {
			PrnLog.Debugf("[%v]quit @ %v",
				enterFlag1, time.Now().Format(TIMEFORMAT_HHMMssMS))

			atomic.StoreUint32(&this.connFlag, 0) // idle
		}()

		PrnLog.Debugf("[%v]enter @ %v",
			enterFlag1, time.Now().Format(TIMEFORMAT_HHMMssMS))

		// ================ loop conn & keepalive ================ //
		for {

			var (
				nano1, wait1 int64
				rttnano      int64
				err          error

				rttMax, rttMin int64
			)
			// ================ conn & retry in connWaitSec ================ //
			wait1 = connWaitSec * NanoTi64          // ns 1000000000
			nano1 = time.Now().UnixNano() - wait1*2 // 第一次立即开始连接
			for true {
				if this.request == nil {
					time.Sleep(time.Second * 1)
					continue
				}

				time.Sleep(time.Millisecond * 200)

				if skipConn { // debug
					PrnLog.Debugf("[%v]no conn in debug @ %v",
						enterFlag1, time.Now().Format(TIMEFORMAT_HHMMssMS))
					break
				}

				if time.Now().UnixNano()-nano1 < wait1 {
					continue
				}
				nano1 = time.Now().UnixNano()

				PrnLog.Infof("[%v]connecting to %v @ %v",
					enterFlag1, this.request.ConnAddr,
					time.Now().Format(TIMEFORMAT_HHMMssMS))

				rttnano, err = this.request.Connect()
				if printDetail {
					PrnLog.Debugf("[%v]connected to %v, rtt=%v @ %v",
						enterFlag1,
						this.request.ConnAddr, libf.NanoToTimeStr(rttnano, 6),
						time.Now().Format(TIMEFORMAT_HHMMssMS))
				}
				if err != nil || this.request.Conn == nil {
					PrnLog.Errorf("[%v]Connect err=%v @ %v",
						enterFlag1, err,
						time.Now().Format(TIMEFORMAT_HHMMssMS))
					// retry conn
					continue
				}

				// ================ InitConnect ================ //
				rttnano, err = this.request.sendInitConnect(enterFlag1)
				if printDetail {
					PrnLog.Debugf("[%v]sendInitConnect, rtt=%v @ %v",
						enterFlag1, libf.NanoToTimeStr(rttnano, 6),
						time.Now().Format(TIMEFORMAT_HHMMssMS))
				}
				if err != nil {
					PrnLog.Errorf("[%v]sendInitConnect err=%v @ %v",
						enterFlag1, err,
						time.Now().Format(TIMEFORMAT_HHMMssMS))
					continue
				}

				break
			}

			if debugflag_conn {
				block := make(chan bool)
				<-block
			}

			// ================ send KeepAlive in KeepAliveSec ================ //
			wait1 = KeepAliveSec * NanoTi64 // ns 1000000000
			nano1 = time.Now().UnixNano()
			rttMax, rttnano, rttMin = 0, 0, 0
			hbCnt := 0
			for true {
				if this.request == nil || this.request.Conn == nil {
					break
				}

				time.Sleep(time.Millisecond * 200)

				if time.Now().UnixNano()-nano1 < wait1 {
					continue
				}
				nano1 = time.Now().UnixNano()

				if skipConn {
					continue
				}

				rttnano, err = this.request.sendkeepAlive(enterFlag1,
					SysConfig.PrintLog.KeepAliveDetail)
				if SysConfig.PrintLog.KeepAliveDetail || debugflag_keepalive {
					PrnLog.Debugf("[%v]sendkeepAlive, rtt=%v @ %v",
						enterFlag1, libf.NanoToTimeStr(rttnano, 6),
						time.Now().Format(TIMEFORMAT_HHMMssMS))
				}
				hbCnt++

				if (0 < rttnano && rttnano < rttMin) || rttMin == 0 {
					rttMin = rttnano
				}
				if rttnano > rttMax {
					rttMax = rttnano
				}
				if !SysConfig.PrintLog.KeepAliveDetail {
					if CheckTs.Check("KeepAlive.Request", 600) {
						PrnLog.Debugf("[%v]send KeepAlive %v, rtt cur=%v,min=%v,max=%v @ %v",
							enterFlag1, hbCnt,
							libf.NanoToTimeStr(rttnano, 6), libf.NanoToTimeStr(rttMin, 6), libf.NanoToTimeStr(rttMax, 6),
							time.Now().Format(TIMEFORMAT_HHMMssMS),
						)
					}
				}

				if err != nil {
					PrnLog.Errorf("[%v]sendkeepAlive err=%v @ %v",
						enterFlag1, err,
						time.Now().Format(TIMEFORMAT_HHMMssMS))

					errNum++
				} else {
					errNum = 0
				}
				if errNum > 3 { // 第 4 次错误时退出
					PrnLog.Errorf("[%v]quit sendkeepAlive, err=%v @ %v",
						enterFlag1, err, time.Now().Format(TIMEFORMAT_HHMMssMS))
					break
				}
			}
		}

		// return
	}(enterFlagIn)

	return
}

// ================ recv from response table ================ /
func RecvSync(SerialNo uint32, toSec int) (pack *FutuPackStu) {
	ts1 := time.Now()
	if toSec <= 0 {
		toSec = 10
	}

	if debugflag_respmng {
		PrnLog.Debugf("[recv].begin, SerialNo=%v, toSec=%v, %v",
			SerialNo, toSec, time.Now().Format("15:04:05.000000"))
	}

	ch := RespPackMng.ChanGet(SerialNo)
	if ch != nil {
		if debugflag_respmng {
			PrnLog.Debugf("[chan]get, ch=%p, %v", ch, time.Now().Format("15:04:05.000000"))
		}
		t := time.NewTimer(time.Duration(toSec) * time.Second)
		select {
		case <-(*ch):
			pack = RespPackMng.PackRead(SerialNo)
		case <-t.C:
			// Retry.
		}
	} else {
		for time.Since(ts1) <= time.Second*time.Duration(toSec) {
			pack = RespPackMng.PackRead(SerialNo)
			if pack != nil {
				break
			}
			time.Sleep(time.Millisecond * 50)
		}
	}
	if debugflag_respmng {
		PrnLog.Debugf("[recv].end, SerialNo=%v, toSec=%v, rtt=%v, %v",
			SerialNo, toSec,
			libf.NanoToTimeStr(time.Now().UnixNano()-ts1.UnixNano(), 6),
			time.Now().Format("15:04:05.000000"))
	}
	RespPackMng.ChanDel(SerialNo)

	if debugflag_respmng {
		PrnLog.Debugf("recv=%v, %v", pack != nil, time.Now().Format("15:04:05.000000"))
	}

	return
}

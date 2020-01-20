package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	SerialNo    uint32
	RespPackMng ResponsePackMngStu
)

/*1
message Request
{
    required bytes HeaderFlag   = 1; // 包头起始标志，固定为“FT”  u8_t szHeaderFlag[2];
    required uint32 ProtoID	    = 2; // 协议ID  u32_t nProtoID;
    required bytes ProtoFmtType = 3; // 协议格式类型，0为Protobuf格式，1为Json格式 u8_t nProtoFmtType;
    required bytes ProtoVer     = 4; // 协议版本，用于迭代兼容, 目前填0 u8_t nProtoVer;
    required uint32 SerialNo    = 5; // 包序列号，用于对应请求包和回包, 要求递增 u32_t nSerialNo;
    required uint32 BodyLen     = 6; // 包体长度  u32_t nBodyLen;
    required bytes BodySHA1     = 7; // 包体原始数据(解密后)的SHA1哈希值  u8_t arrBodySHA1[20];
    required bytes Reserved     = 8; // 保留8字节扩展  u8_t arrReserved[8];
}*/
type FutuPackStu struct {
	HeaderFlag   [2]uint8  // u8_t szHeaderFlag[2];
	ProtoID      uint32    // u32_t nProtoID;
	ProtoFmtType uint8     // u8_t nProtoFmtType;
	ProtoVer     uint8     // u8_t nProtoVer;
	SerialNo     uint32    // u32_t nSerialNo;
	BodyLen      uint32    // u32_t nBodyLen;
	BodySHA1     [20]uint8 // u8_t arrBodySHA1[20];
	Reserved     [8]uint8  // u8_t arrReserved[8];
	body         []byte    // []byte add;
}

// SetProtoID set nProtoID
func (this *FutuPackStu) ProtoIDSet(nProtoID uint32) {
	this.ProtoID = nProtoID
}

// SetBody set body
func (this *FutuPackStu) BodySet(body []byte) {
	this.body = body
	this.BodyLen = uint32(len(body))

	sha := sha1.New()
	sha.Write(this.body)
	arrBodySHA1 := sha.Sum(nil)
	copy(this.BodySHA1[:], arrBodySHA1)
}

func (this *FutuPackStu) SerialNoStrGet() string {
	return fmt.Sprintf("p_%d", this.SerialNo)
}

// Pack pack
func (this *FutuPackStu) PackBuild() ([]byte, error) {
	var err error

	// var rnd int = libf.GetRand(1000) // debug

	this.HeaderFlag = [2]byte{'F', 'T'}
	this.ProtoFmtType = uint8(0)
	this.ProtoVer = 0
	this.SerialNo = atomic.AddUint32(&SerialNo, 1) // debugatomic.AddUint32(&SerialNo, uint32(rnd)) //
	// var arrReservedTmp [8]uint8
	// copy(this.Reserved[:], arrReservedTmp[:8])

	packBuf := new(bytes.Buffer)
	err = binary.Write(packBuf, binary.LittleEndian, &this.HeaderFlag)   // PackBuild()
	err = binary.Write(packBuf, binary.LittleEndian, &this.ProtoID)      // ProtoIDSet()
	err = binary.Write(packBuf, binary.LittleEndian, &this.ProtoFmtType) // PackBuild()
	err = binary.Write(packBuf, binary.LittleEndian, &this.ProtoVer)     // PackBuild()
	err = binary.Write(packBuf, binary.LittleEndian, &this.SerialNo)     // PackBuild()
	err = binary.Write(packBuf, binary.LittleEndian, &this.BodyLen)      // BodySet()
	err = binary.Write(packBuf, binary.LittleEndian, &this.BodySHA1)     // BodySet()
	err = binary.Write(packBuf, binary.LittleEndian, &this.Reserved)

	err = binary.Write(packBuf, binary.LittleEndian, &this.body)

	return packBuf.Bytes(), err
}

// PackRead from arrPack to this FutuPackStu
func (this *FutuPackStu) PackRead(arrPack []byte) error {
	var err error
	reader := bytes.NewReader(arrPack)
	err = binary.Read(reader, binary.LittleEndian, &this.HeaderFlag)
	err = binary.Read(reader, binary.LittleEndian, &this.ProtoID)
	err = binary.Read(reader, binary.LittleEndian, &this.ProtoFmtType)
	err = binary.Read(reader, binary.LittleEndian, &this.ProtoVer)
	err = binary.Read(reader, binary.LittleEndian, &this.SerialNo)
	err = binary.Read(reader, binary.LittleEndian, &this.BodyLen)
	err = binary.Read(reader, binary.LittleEndian, &this.BodySHA1)
	err = binary.Read(reader, binary.LittleEndian, &this.Reserved)

	this.body = make([]byte, this.BodyLen)
	err = binary.Read(reader, binary.LittleEndian, &this.body)

	return err
}

// to string
func (this *FutuPackStu) String() string {
	return fmt.Sprintf("BodyLen: %d body: %s",
		this.BodyLen,
		this.body,
	)
}

// GetBody get body data
func (this *FutuPackStu) BodyGet() []byte {
	return this.body
}

type ResponsePackMngStu struct {
	packList []*FutuPackStu
	ts       []int64
	seqmap   map[uint32]int // serial no -- seq

	chanmap map[uint32](*chan bool) // serial no -- chan

	sync.RWMutex
}

func (this *ResponsePackMngStu) Init() {
	this.packList = make([]*FutuPackStu, 0)
	this.ts = make([]int64, 0)
	this.seqmap = make(map[uint32]int, 0)

	this.chanmap = make(map[uint32](*(chan bool)), 0)

	return
}

func (this *ResponsePackMngStu) ChanAdd(SerialNoIn uint32) {
	ch := make(chan bool, 1) // 不提供数量的情况下写入时堵塞， new(chan bool)
	defer func() {
		if debugflag_respmng {
			PrnLog.Debugf("[chan]add, ch=%p, SerialNo=%v, %v",
				&ch, SerialNoIn, time.Now().Format("15:04:05.000000"))
		}
	}()

	this.Lock()
	this.chanmap[SerialNoIn] = &ch
	this.Unlock()
	return
}

func (this *ResponsePackMngStu) ChanGet(SerialNoIn uint32) *(chan bool) {
	var ch *(chan bool)
	defer func() {
		if debugflag_respmng {
			PrnLog.Debugf("[chan]get, ch=%p, SerialNo=%v, %v",
				ch, SerialNoIn, time.Now().Format("15:04:05.000000"))
		}
	}()
	this.RLock()
	ch = this.chanmap[SerialNoIn]
	this.RUnlock()
	return ch
}

func (this *ResponsePackMngStu) ChanDel(SerialNoIn uint32) {
	var (
		ch *(chan bool)
		ok bool
	)
	defer func() {
		if debugflag_respmng {
			PrnLog.Debugf("[chan]del, ch=%p, SerialNo=%v, ok=%v, %v",
				ch, SerialNoIn, ok, time.Now().Format("15:04:05.000000"))
		}
	}()
	this.Lock()
	ch, ok = this.chanmap[SerialNoIn]
	if ok {
		if ch != nil {
			close(*ch)
		}
		delete(this.chanmap, SerialNoIn)
	}
	this.Unlock()
	return
}

func (this *ResponsePackMngStu) PackAdd(packIn *FutuPackStu) {

	var (
		seq int
		ok  bool
	)

	this.Lock()
	defer func() {
		this.Unlock()
		ch := this.ChanGet(packIn.SerialNo)
		if debugflag_respmng {
			PrnLog.Debugf("[chan]get, ch=%p, SerialNo=%v, seq=%v, %v",
				ch, packIn.SerialNo, seq, time.Now().Format("15:04:05.000000"))
		}
		if ch != nil {
			(*ch) <- true
			if debugflag_respmng {
				PrnLog.Debugf("[chan]get, ch=%p, set done, %v ",
					ch, time.Now().Format("15:04:05.000000"))
			}
		}
	}()

	if seq, ok = this.seqmap[packIn.SerialNo]; ok {
		this.packList[seq] = packIn
		this.ts[seq] = time.Now().Unix()
		return
	}
	this.packList = append(this.packList, packIn)
	this.ts = append(this.ts, time.Now().Unix())
	seq = len(this.packList) - 1
	this.seqmap[packIn.SerialNo] = seq

	return
}

func (this *ResponsePackMngStu) PackRead(SerialNoIn uint32) (packOut *FutuPackStu) {

	this.RLock()

	if seq, ok := this.seqmap[SerialNoIn]; ok {

		len1 := len(this.packList)
		isValid := 0 <= seq && seq < len1

		if !isValid {
			this.RUnlock()
			return
		}

		packOut = this.packList[seq]

		this.RUnlock()

		this.Lock()
		this.packList = append(this.packList[:seq], this.packList[seq+1:]...)
		this.ts = append(this.ts[:seq], this.ts[seq+1:]...)
		this.resetMap()
		this.Unlock()

		return
	}
	this.RUnlock()

	return
}

// 删除老化的记录，建议 3-5 分钟扫描一次
// toSec 超时秒数,建议 300-600 秒（5-10分钟）
func (this *ResponsePackMngStu) Scan(toSec int64) {

	if toSec <= 0 {
		toSec = 300
	}

	this.Lock()

	for seq := 0; seq < len(this.ts); {
		if time.Now().Unix()-this.ts[seq] > toSec {
			SerialNo := this.packList[seq].SerialNo
			this.packList = append(this.packList[:seq], this.packList[seq+1:]...)
			this.ts = append(this.ts[:seq], this.ts[seq+1:]...)

			if ch, ok := this.chanmap[SerialNo]; ok {
				if ch != nil {
					close(*ch)
				}
				delete(this.chanmap, SerialNo)
			}
			continue
		}
		seq++
	}
	this.resetMap()
	this.Unlock()

	return
}

// 调用者负责lock
func (this *ResponsePackMngStu) resetMap() {
	this.seqmap = make(map[uint32]int, 0)
	for seq := 0; seq < len(this.packList); seq++ {
		this.seqmap[this.packList[seq].SerialNo] = seq
	}
}

package main

import (
	"errors"
	"fmt"
	"time"

	GetGlobalState "github.com/futuopen/ftapi4go/pb/getglobalstate"
	Qot_Common "github.com/futuopen/ftapi4go/pb/qotcommon"
	Qot_GetBasicQot "github.com/futuopen/ftapi4go/pb/qotgetbasicqot"
	Qot_GetSecuritySnapshot "github.com/futuopen/ftapi4go/pb/qotgetsecuritysnapshot"
	Qot_Sub "github.com/futuopen/ftapi4go/pb/qotsub"

	libf "libfunc"

	"github.com/gogo/protobuf/proto"
)

var futu_cli_debugflag = false

func fQot_GetBasicQot_debug(Request *RequestStu) {

	var securityList []*Qot_Common.Security //// modify ////
	// securityList = append(securityList, &Qot_Common.Security{
	// 	Market: int32(11),
	// 	Code:   "AAPL",
	// })
	securityList = append(securityList, &Qot_Common.Security{
		Market: int32(1),
		Code:   "01846",
	})
	securityList = append(securityList, &Qot_Common.Security{
		Market: int32(1),
		Code:   "01830",
	})

	// 订阅数据类型参考 file:///E:/futu/docs/%E5%AF%8C%E9%80%94API%E8%AF%B4%E6%98%8E%E6%96%87%E6%A1%A3/protocol/base_define.html#subtype
	var subTypeList []int32 // 订阅多只股票不用写多个类型？
	//subTypeList = append(subTypeList, []int32{1, 1}...)
	// subTypeList = append(subTypeList, []int32{1}...)
	subTypeList = append(subTypeList, int32(1))
	// isRegOrUnRegPush := true

	// 复权类型参考 file:///E:/futu/docs/%E5%AF%8C%E9%80%94API%E8%AF%B4%E6%98%8E%E6%96%87%E6%A1%A3/protocol/base_define.html#rehabtype-k
	var regPushRehabTypeList []int32
	// regPushRehabTypeList = append(regPushRehabTypeList, []int32{1, 1}...)
	// regPushRehabTypeList = append(regPushRehabTypeList, []int32{1}...)
	regPushRehabTypeList = append(regPushRehabTypeList, int32(1))

	fQot_Sub(Request,
		Qot_Sub.C2S{ //// modify ////
			SecurityList:         securityList,
			SubTypeList:          subTypeList,
			IsSubOrUnSub:         true, //ture表示订阅,false表示反订阅
			IsRegOrUnRegPush:     true, //是否注册或反注册该连接上面行情的推送,该参数不指定不做注册反注册操作
			RegPushRehabTypeList: regPushRehabTypeList,
			// IsUnsubAll:false,//一键取消当前连接的所有订阅,当被设置为True时忽略其他参数。
		}, true,
	)

	fQot_GetBasicQot(Request,
		Qot_GetBasicQot.C2S{ //// modify ////
			SecurityList: securityList,
		}, true,
	)

	return
}

func fQot_Sub(Request *RequestStu, c2s Qot_Sub.C2S, prndetailIn ...bool) (futResp Qot_Sub.Response) {

	var (
		ProtoID = uint32(3001) //// modify ////
		n       int
		rttnano int64
	)
	prndetail := false
	if len(prndetailIn) > 0 {
		prndetail = prndetailIn[0]
	}

	// pack
	pack := &FutuPackStu{}
	pack.ProtoIDSet(ProtoID)

	futuReq := &Qot_Sub.Request{ //// modify ////
		C2S: &c2s,
	}

	pbData, err := proto.Marshal(futuReq)
	if err != nil {
		return
	}
	pack.BodySet(pbData)

	var packResp *FutuPackStu
	packResp, rttnano, n, err = Request.Send(pack, true) //// modify true or false ////
	if err != nil || futu_cli_debugflag {
		PrnLog.Debugf("Send n=%v, err=%v, rtt=%v @ %v", n, err,
			libf.NanoToTimeStr(rttnano, 3), time.Now().Format("15:04:05.000000"))
		if err != nil {
			PrnLog.Errorf("Send error:%v", err)
			return
		}
	}

	if packResp.ProtoID != ProtoID {
		PrnLog.Warningf("packResp.ProtoID=%v, send ProtoID=%v", packResp.ProtoID, ProtoID)
		return
	}

	if futu_cli_debugflag {
		PrnLog.Debugf("%v==> @ %v", packResp.ProtoID, time.Now().Format("15:04:05.000000"))
	}

	fut := &Qot_Sub.Response{} //// modify ////
	err = proto.Unmarshal(packResp.body, fut)
	futResp = *fut
	if err != nil {
		PrnLog.Errorf("unmarshaling error:%v", err)
		return
	}

	if fut.GetRetType() != 0 {
		PrnLog.Errorf("@ %v, Ret=%v, Msg=%v ", time.Now().Format("15:04:05.000000"),
			fut.GetRetType(), fut.GetRetMsg(),
		)
		return
	}

	if prndetail {
		PrnLog.Debugf("rtt=%v, @ %v, resp=%v ", libf.NanoToTimeStr(rttnano, 3), time.Now().Format("15:04:05.000000"),
			fut.String(),
		)
	}

	return
}

/* 先订阅 fQot_Sub(Request *RequestStu) 再查询。每次启动程序(重新连接)都要订阅 */
func fQot_GetBasicQot(Request *RequestStu,
	c2s Qot_GetBasicQot.C2S, prndetailIn ...bool) (futResp Qot_GetBasicQot.Response) {

	var (
		ProtoID = uint32(3004) //// modify ////
		n       int
		rttnano int64
	)
	prndetail := false
	if len(prndetailIn) > 0 {
		prndetail = prndetailIn[0]
	}

	// pack
	pack := &FutuPackStu{}
	pack.ProtoIDSet(ProtoID)

	futuReq := &Qot_GetBasicQot.Request{ //// modify ////
		C2S: &c2s,
	}

	pbData, err := proto.Marshal(futuReq)
	if err != nil {
		return
	}
	pack.BodySet(pbData)

	var packResp *FutuPackStu
	packResp, rttnano, n, err = Request.Send(pack, true) //// modify true or false ////
	if err != nil || futu_cli_debugflag {
		PrnLog.Debugf("Send n=%v, err=%v, rtt=%v @ %v", n, err,
			libf.NanoToTimeStr(rttnano, 3), time.Now().Format("15:04:05.000000"))
		if err != nil {
			PrnLog.Errorf("Send error:%v", err)
			return
		}
	}

	if packResp.ProtoID != ProtoID {
		PrnLog.Warningf("packResp.ProtoID=%v, send ProtoID=%v", packResp.ProtoID, ProtoID)
		return
	}

	if futu_cli_debugflag {
		PrnLog.Debugf("%v==> @ %v", packResp.ProtoID, time.Now().Format("15:04:05.000000"))
	}

	fut := &Qot_GetBasicQot.Response{} //// modify ////
	err = proto.Unmarshal(packResp.body, fut)
	futResp = *fut
	if err != nil {
		PrnLog.Errorf("unmarshaling error:%v", err)
		return
	}

	if fut.GetRetType() != 0 {
		PrnLog.Errorf("@ %v, Ret=%v, Msg=%v ", time.Now().Format("15:04:05.000000"),
			fut.GetRetType(), fut.GetRetMsg(),
		)
		return
	}

	PrnLog.Debugf("rtt=%v, num=%v, resp: ", libf.NanoToTimeStr(rttnano, 3), len(fut.S2C.BasicQotList))
	if prndetail {
		/*1 see .\readme\9.fQot_GetBasicQot.txt */
		for seq1, bs1 := range fut.S2C.BasicQotList {
			s1 := bs1.GetSecurity()
			// , market=%v, SecType=%v,
			// , bs1.GetName(), s1.GetMarket(), bs1.GetSecType(),GetAmplitude
			fmt.Printf("seq=%v, code=%v, list=%v, upd=%v, O=%v, H=%v, L=%v, C=%v, LC=%v, vol=%0.2f, amo=%0.2f, rate=%v%%\n",
				seq1, s1.GetCode(), bs1.GetListTime(), bs1.GetUpdateTime(),
				bs1.GetOpenPrice(), bs1.GetHighPrice(), bs1.GetLowPrice(), bs1.GetCurPrice(),
				bs1.GetLastClosePrice(),
				float64(bs1.GetVolume())/10000.0, float64(bs1.GetTurnover())/10000.0,
				bs1.GetTurnoverRate(),
			)
		}
	}

	if false {
		PrnLog.Debugf("rtt=%v, @ %v, resp=%v ", libf.NanoToTimeStr(rttnano, 3), time.Now().Format("15:04:05.000000"),
			fut.String(),
		)
	}

	return
}

func fQot_GetSecuritySnapshot_debug(Request *RequestStu) {

	var securityList []*Qot_Common.Security //// modify ////
	// 美股不支持行情？
	if false {
		securityList = append(securityList, &Qot_Common.Security{
			Market: int32(11),
			Code:   "AAPL",
		})
	}
	if false {
		securityList = append(securityList, &Qot_Common.Security{
			Market: int32(1),
			Code:   "01846",
		})
		securityList = append(securityList, &Qot_Common.Security{
			Market: int32(1),
			Code:   "01830",
		})
	}
	if true {
		securityList = append(securityList, &Qot_Common.Security{
			Market: int32(1),
			Code:   "00007",
		})
		securityList = append(securityList, &Qot_Common.Security{
			Market: int32(1),
			Code:   "00150",
		})

		securityList = append(securityList, &Qot_Common.Security{
			Market: int32(1),
			Code:   "07230",
		})
		// securityList = append(securityList, &Qot_Common.Security{
		// 	Market: int32(1),
		// 	Code:   "09146",
		// })
	}
	for {
		_, err := fQot_GetSecuritySnapshot(Request,
			Qot_GetSecuritySnapshot.C2S{
				SecurityList: securityList,
			}, true)
		if err != nil {
			PrnLog.Errorf("err=%v", err)
		}
		time.Sleep(time.Second * 3)
	}

	// return
}

/*1 .\readme\6.futu权限.txt
获取股票快照
    请求协议ID: 3203
    每次最多可请求股票数与用户等级相关, 一级: 400 , 二级: 300 , 三级: 200
    30秒内快照最多请求次数与用户等级相关，一级: 30 , 二级: 20 , 三级: 10
*/
func fQot_GetSecuritySnapshot(Request *RequestStu, c2s Qot_GetSecuritySnapshot.C2S,
	prndetailIn ...bool) (futResp Qot_GetSecuritySnapshot.Response, err error) {

	var (
		ProtoID = uint32(3203) //// modify ////
		n       int
		rttnano int64

		securityLen = len(c2s.SecurityList)
	)
	prndetail := false
	if len(prndetailIn) > 0 {
		prndetail = prndetailIn[0]
	}

	if securityLen == 0 {
		err = errors.New("no c2s.SecurityList ")
		return
	}

	// pack
	pack := &FutuPackStu{}
	pack.ProtoIDSet(ProtoID)

	futuReq := &Qot_GetSecuritySnapshot.Request{ //// modify ////
		C2S: &c2s,
	}

	if futu_cli_debugflag {
		PrnLog.Debugf("SecurityList len=%v, %v...%v @ %v",
			len(c2s.SecurityList), c2s.SecurityList[0].Code, c2s.SecurityList[securityLen-1].Code,
			time.Now().Format("15:04:05.000000"))
	}

	pbData, err2 := proto.Marshal(futuReq)
	if err2 != nil {
		err = fmt.Errorf("Marshal Request err=%v", err2)
		return
	}
	pack.BodySet(pbData)

	var packResp *FutuPackStu
	packResp, rttnano, n, err = Request.Send(pack, true)
	if err != nil || futu_cli_debugflag {
		PrnLog.Debugf("Send n=%v, err=%v, rtt=%v @ %v", n, err,
			libf.NanoToTimeStr(rttnano, 3), time.Now().Format("15:04:05.000000"))
		if err != nil {
			PrnLog.Errorf("Send error:%v", err)
			err = fmt.Errorf("Send err=%v", err)
			return
		}
	}

	if packResp.ProtoID != ProtoID {
		PrnLog.Warningf("packResp.ProtoID=%v, send ProtoID=%v", packResp.ProtoID, ProtoID)
		return
	}

	if futu_cli_debugflag {
		PrnLog.Debugf("%v==> @ %v", packResp.ProtoID, time.Now().Format("15:04:05.000000"))
	}

	fut := &Qot_GetSecuritySnapshot.Response{} //// modify ////
	err = proto.Unmarshal(packResp.body, fut)
	futResp = *fut
	if err != nil {
		PrnLog.Errorf("unmarshaling error:%v", err)
		err = fmt.Errorf("Marshal packResp err=%v", err)
		return
	}

	if fut.GetRetType() != 0 {
		PrnLog.Errorf("@ %v, Ret=%v, Msg=%v ", time.Now().Format("15:04:05.000000"),
			fut.GetRetType(), fut.GetRetMsg(),
		)
		err = fmt.Errorf("fut.GetRetType()=%v,msg=%v", fut.GetRetType(), fut.GetRetMsg())
		return
	}

	if prndetail {
		PrnLog.Debugf("rtt=%v, @ %v, resp: ", libf.NanoToTimeStr(rttnano, 3), time.Now().Format("15:04:05.000000"))
		for seq1, v := range fut.S2C.SnapshotList {
			bs1 := v.Basic
			// sinfo := v.Basic.Security.Code
			fmt.Printf("seq=%v, code=%v, list=%v, upd=%v, O=%v, H=%v, L=%v, C=%v, LC=%v, vol=%0.2f, amo=%0.2f, rate=%v%%\n",
				seq1, v.Basic.Security.Code, bs1.GetListTime(), bs1.GetUpdateTime(),
				bs1.GetOpenPrice(), bs1.GetHighPrice(), bs1.GetLowPrice(), bs1.GetCurPrice(),
				bs1.GetLastClosePrice(),
				float64(bs1.GetVolume())/10000.0, float64(bs1.GetTurnover())/10000.0,
				bs1.GetTurnoverRate(),
			)
		}
	}

	return
}

func fGetGlobalState(Request *RequestStu) {

	var (
		ProtoID = uint32(1002)
		n       int
		rttnano int64
	)

	// pack
	pack := &FutuPackStu{}
	pack.ProtoIDSet(ProtoID)
	futuReq := &GetGlobalState.Request{
		C2S: &GetGlobalState.C2S{
			UserID: Request.FutuID,
		},
	}

	pbData, err := proto.Marshal(futuReq)
	if err != nil {
		return
	}
	pack.BodySet(pbData)

	var packResp *FutuPackStu
	packResp, rttnano, n, err = Request.Send(pack, true)
	if err != nil || futu_cli_debugflag {
		PrnLog.Debugf("Send n=%v, err=%v, rtt=%v @ %v", n, err,
			libf.NanoToTimeStr(rttnano, 3), time.Now().Format("15:04:05.000000"))
		if err != nil {
			PrnLog.Errorf("Send error:%v", err)
			return
		}
	}

	if packResp.ProtoID != ProtoID {
		PrnLog.Warningf("packResp.ProtoID=%v, send ProtoID=%v", packResp.ProtoID, ProtoID)
		return
	}

	if futu_cli_debugflag {
		PrnLog.Debugf("%v==> @ %v", packResp.ProtoID, time.Now().Format("15:04:05.000000"))
	}

	fut := &GetGlobalState.Response{}
	err = proto.Unmarshal(packResp.body, fut)
	if err != nil {
		PrnLog.Errorf("unmarshaling error:%v", err)
		return
	}

	if fut.GetRetType() != 0 {
		PrnLog.Errorf("@ %v, Ret=%v, Msg=%v ", time.Now().Format("15:04:05.000000"),
			fut.GetRetType(), fut.GetRetMsg(),
		)
		return
	}

	PrnLog.Debugf("@ %v, resp=%v ", time.Now().Format("15:04:05.000000"), fut.String())

	return
}

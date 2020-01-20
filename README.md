
## Example
detail see ./readme/*

```
package main

import (
	"testing"
)

func Test_cli(t *testing.T) {
	var Request RequestStu       // 定义一个连接实体
	var ConnManage ConnManageStu // 定义一个连接管理实体，里边自动管理连接、init、维持心跳、重新连接

	CommonEnter() // 公共初始化

	Request.ConnAddr = SysConfig.ConnAddr // 设置地址
	ConnManage.Enter(&Request, true)      // 启动连接管理
	Request.ReadyWait(0)                  //连接就绪之后再往下

	// ========================= 服务测试
	// fGetGlobalState(&Request)
	// fQot_GetBasicQot_debug(&Request)

	fQot_GetSecuritySnapshot_debug(&Request)

	return

	// wait
	block := make(chan bool)
	<-block
}


```
```
log example
 [01-06 17:16:07]|I|futu_cli|debug.go|init()|8| invoke
 [01-06 17:16:07]|I|futu_cli|config.go|Init_1Config()|13| invoke
 [01-06 17:16:07]|I|futu_cli|config.go|Load()|38| begin ......
 [01-06 17:16:07]|I|futu_cli|config.go|Load()|58| end...
    config={"ConnAddr":"127.0.0.1:11111","LogLevel":"DEBUG","KeepAlive":5,"PrintLog":{"KeepAliveDetail":false}}
 [01-06 17:16:07]|I|futu_cli|value.go|Init_Value()|29| invoke
 [01-06 17:16:07]|I|futu_cli|server_com.go|Init_server_common()|8| invoke
 [01-06 17:16:07]|I|futu_cli|server_com.go|CommonEnter()|29| Start success.
 [01-06 17:16:07]|D|futu_cli|1socket.go|Enter()|385| [124]enter @ 17:16:07.284863
 [01-06 17:16:07]|I|futu_cli|1socket.go|Enter()|420| [124]connecting to 127.0.0.1:11111 @ 17:16:07.486578
 [01-06 17:16:07]|D|futu_cli|1socket.go|Enter()|426| [124]connected to 127.0.0.1:11111, rtt=0.007848 @ 17:16:07.494426
 [01-06 17:16:07]|D|futu_cli|1socket.go|sendInitConnect()|156| @ 17:16:07.496432, resp=retType:0 retMsg:"" errCode:0 
    s2c:<serverVer:208 loginUserID:**** connID:***** connAESKey:"****" keepAliveInterval:10 aesCBCiv:"****" >
 [01-06 17:16:07]|D|futu_cli|1socket.go|Enter()|441| [124]sendInitConnect, rtt=0.011029 @ 17:16:07.505456
 [01-06 17:16:07]|D|futu_cli|main.go|dbg_GetGlobalState()|57| Send n=51, err=<nil> @ 17:16:07.686970
 [01-06 17:16:07]|D|futu_cli|main.go|dbg_GetGlobalState()|65| 1002==> @ 17:16:07.686970
 [01-06 17:16:07]|D|futu_cli|main.go|dbg_GetGlobalState()|74| @ 17:16:07.686970, resp=retType:0 retMsg:"" errCode:0 
    s2c:<marketHK:6 marketUS:8 marketSH:6 marketSZ:6 marketHKFuture:13 qotLogined:true trdLogined:true serverVer:208 serverBuildNo:700 
    time:1578302167 localTime:1.57830216768697e+09 programStatus:<type:ProgramStatusType_Ready strExtDesc:"" > 
    qotSvrIpAddr:"115.159.18.59" trdSvrIpAddr:"115.159.18.59" >
 [01-06 17:16:12]|D|futu_cli|1socket.go|Enter()|502| [124]send KeepAlive 1, rtt cur=0.001001,min=0.001001,max=0.001001 @ 17:16:12.519922
```

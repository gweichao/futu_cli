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

func _Test_RequestStu(t *testing.T) {

	var Request RequestStu
	var HeartBeat ConnManageStu

	CommonEnter()
	Request.ConnAddr = SysConfig.ConnAddr

	if true {
		HeartBeat.Enter(&Request, true)
		// HeartBeat.Enter(&Request, true, true)
		// time.Sleep(time.Second * 20)
		// HeartBeat.routineQuit(0, true, true)
		// time.Sleep(time.Second * 5)
		// HeartBeat.Enter(&Request, true, true)
		// time.Sleep(time.Second * 20)
	}

	if false {
		Request.Close()
	}

	// wait
	block := make(chan bool)
	<-block

	return
}

package main

import (
	"runtime"
)

func Init_server_common() {
	PrnLog.Infof("invoke ")

	nc := runtime.NumCPU()
	if nc >= 2 {
		nc--
	}
	runtime.GOMAXPROCS(nc)

}

func CommonEnter4Test() {
	CommonEnter()
}
func CommonEnter() {

	Init_1Config()

	Init_Value()

	Init_server_common()

	PrnLog.Infof("Start success.")

}

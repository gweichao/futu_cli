package main

import (
	libf "libfunc"
)

func init() {
	PrnLog.Infof("invoke ")
}

var PrnLog *libf.PrnLogStu = libf.PrnLogGet(SysFlag)

var Errorf = func(fmt string, args ...interface{}) {
	PrnLog.LogPrint(libf.LOG_ERROR, 1, false, false, fmt, args...)
}

var Debugf = func(fmt string, args ...interface{}) {
	PrnLog.LogPrint(libf.LOG_DEBUG, 1, false, false, fmt, args...)
}

var Infof = func(fmt string, args ...interface{}) {
	PrnLog.LogPrint(libf.LOG_INFO, 1, false, false, fmt, args...)
}
var Warningf = func(fmt string, args ...interface{}) {
	PrnLog.LogPrint(libf.LOG_WARNING, 1, false, false, fmt, args...)
}

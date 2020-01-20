package main

import (

	// "sync"

	libf "libfunc"
)

const (
	ClientVer int32 = 317                // 客户端版本号，clientVer = "."以前的数 * 100 + "."以后的，举例：2.21版本为2 * 100 + 21 = 221
	ClientId        = "qzsxhnbpjyfpmqd4" // 客户端唯一标识，无生具体生成规则，客户端自己保证唯一性即可
)
const NanoTf64 float64 = 1000000000.0
const NanoTi64 int64 = 1000000000

const (
	DataSizeDefault             int    = 4096
	TIMEFORMAT_YYYYMMDDHHMMssMS string = "2006-01-02 15:04:05.000000"
	TIMEFORMAT_HHMMssMS         string = "15:04:05.000000"
	// 要打印秒后面的时间，+若干个9就行，比如 .999999 打印到6位
	FormatYMDHMSZone = "2006-01-02 15:04:05 -0700"
	FormatYMDHMS     = "2006-01-02 15:04:05"
	FormatDTimeNoDot = "20060102150405"
	FormatYMD        = "2006-01-02"
	FormatHMS        = "15:04:05"

	SysFlag = "futu_cli"
)

var (
	SysConfig SysConfigStu

	// 发送邮件
	// MailFrom, err = MailMng.Sendmail(this.MailTo, this.Subject, mailtype, []byte(Content1))
	MailMng libf.MailStu
)

func Init_Value() {
	PrnLog.Infof("invoke ")

}

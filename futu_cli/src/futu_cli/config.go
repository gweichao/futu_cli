package main

import (
	"os"
	"strings"
	"toml"

	libf "libfunc"
)

func Init_1Config() {

	PrnLog.Infof("invoke ")

	SysConfig.Load("", true)

}

type SysConfigStu struct {
	ConnAddr string // futu opend 的地址端口

	LogLevel string

	KeepAlive int

	PrintLog struct {
		KeepAliveDetail bool
	}

	MailsettingList []libf.MailsettingStu
}

func (this *SysConfigStu) Load(file1 string, prn bool) {

	if len(file1) == 0 {
		file1 = "config.toml"
	}

	if prn {
		PrnLog.Infof(" begin ...... ")
	}

	file1 = libf.FindPathFileInNameList([]string{file1, "futu_cli.toml"})

	defer func() {
		if len(this.MailsettingList) > 0 {
			MailMng.AddSendSetting(this.MailsettingList)
		}
		if this.KeepAlive <= 0 {
			this.KeepAlive = 5
		}

		this.ConnAddr = strings.TrimSpace(this.ConnAddr)
		if len(this.ConnAddr) == 0 {
			this.ConnAddr = "127.0.0.1:11111"
		} else if !strings.Contains(this.ConnAddr, ":") {
			this.ConnAddr = "127.0.0.1:" + this.ConnAddr
		} else if strings.HasPrefix(this.ConnAddr, ":") {
			this.ConnAddr = "127.0.0.1" + this.ConnAddr
		}

		if prn {
			PrnLog.Infof("end...\nconfig=%v", libf.StructToJsonStringOne(this))
		}
	}()

	var err error
	if _, err = toml.DecodeFile(file1, this); err != nil {
		PrnLog.Errorf("error load config file! %v", err)
		os.Exit(-1)
	}

	this.LogLevel = strings.TrimSpace(strings.ToUpper(this.LogLevel))
	if len(this.LogLevel) == 0 {
		this.LogLevel = libf.DEBUGName
	}

	PrnLog.SetLogLevel(this.LogLevel, "")

	return
}

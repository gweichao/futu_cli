package libfunc

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/smtp" // golang官方库net/smtp支持的是starttls, 通过发给25端口，然后在email协议层进行tls密钥协商、加密
	"strings"
	"time"
)

/*e.g.
step1
	// 定义发送管理实体
	var MailMng      libf.MailStu
	// 定义发送设置，可多个，实际发送时按顺序选择
	var Mailsetting []libf.MailsettingStu = []libf.MailsettingStu{
		{SendSvrAddr: "smtp.163.com", SendSvrPort: "994", NetType: "tls", SenderMail: "test<test@test.com>", SenderPassword: "***"},
	}
step2
	// 添加配置
	MailMng.AddSendSetting(Mailsetting)
	// 发送邮件
	this.MailFrom, err = MailMng.Sendmail(this.MailTo, this.Subject, mailtype, []byte(Content1))
step3 在toml文件的配置方法
	结构体定义
		type configStu struct {
			MailsettingList []MailsettingStu
		}
	toml配置方法 -- 放在文件的末尾，至少单个配置项之后
		[[MailsettingList]]
		  SendSvrAddr = "smtp.163.com"
		  SendSvrPort = "994"
		  NetType = "normal"
		  SenderMail = "test<test@test.com>"
		  SenderPassword = "***"

		[[MailsettingList]]
		  SendSvrAddr = "smtp.163.com"
		  SendSvrPort = "25"
		  NetType = "tls"
		  SenderMail = "test<test@163.com>"
		  SenderPassword = "***"
*/

type MailStu struct {
	sendSvrList []sendSvrStu
	charType    int // 0 不处理，1 用utf-8, 2用gbk
}
type sendSvrStu struct {
	netType int // nettype 1 normal; 2 tls
	/* 发送邮件的smtp服务器地址，端口
	举例 "smtp.163.com" "994" */
	sendSvrAddr, sendSvrPort string

	/* 举例 "管理员<test@test.com>"，ttt@qq.com */
	senderMail   string // 发送者邮箱, 填写from字段
	senderPasswd string // 密码 用于登录smtp

	_sendSvr_AddrPort string // 内部使用，地址+端口
	// 拆出来的邮箱、用户名，senderAddr用于登录smtp等。senderName 暂无用途
	_sender_Addr, _sender_Name string

	charType int // 0 不处理，1 用utf-8, 2用gbk
}

func (this *MailStu) SetCharType(charTypeIn int) {
	this.charType = charTypeIn
	for i := 0; i < len(this.sendSvrList); i++ {
		this.sendSvrList[i].SetCharType(charTypeIn)
	}
	return
}

type MailsettingStu struct {
	SendSvrAddr    string
	SendSvrPort    string
	NetType        string // normal; tls
	SenderMail     string
	SenderPassword string
}

func (this *MailStu) AddSendSetting(settingList []MailsettingStu) {
	for _, setting := range settingList {
		this.addSendSvr(setting)
	}
	return
}

/* NetType：1 normal; 2 tls */
func (this *MailStu) addSendSvr(settingList MailsettingStu) {

	var sendSvr sendSvrStu

	if settingList.NetType == "normal" {
		sendSvr.netType = 1
	} else { // if settingList.NetTypeS == "tls" {
		sendSvr.netType = 2
	}

	sendSvr.sendSvrAddr, sendSvr.sendSvrPort = settingList.SendSvrAddr, settingList.SendSvrPort

	sendSvr.senderMail, sendSvr.senderPasswd = settingList.SenderMail, settingList.SenderPassword

	sendSvr._sendSvr_AddrPort = fmt.Sprintf("%s:%s", sendSvr.sendSvrAddr, sendSvr.sendSvrPort)

	strlist1 := strings.Split(sendSvr.senderMail, "<")
	if len(strlist1) > 1 {
		sendSvr._sender_Name = strlist1[0]
		strlist2 := strings.Split(strlist1[1], ">")
		sendSvr._sender_Addr = strlist2[0]
	} else {
		sendSvr._sender_Addr = sendSvr.senderMail
		sendSvr._sender_Name = ""
	}
	sendSvr._sender_Addr = StringTrim3(sendSvr._sender_Addr)

	this.sendSvrList = append(this.sendSvrList, sendSvr)

	//prnLog.Debugf( "%+v", *this)

	return
}

/* mailtype:html plain */
func (this *MailStu) Sendmail(to, subject, mailtype string, body []byte) (sender string, rtnerr error) {

	if len(this.sendSvrList) == 0 {
		rtnerr = errors.New("no mail sender setting")
		return
	}

	if len(body) == 0 {
		body = []byte("我是一封电子邮件!golang发出.tls, " + time.Now().Format("2006-01-02 15:04:05"))
	}

	for i := 0; i < len(this.sendSvrList); i++ {
		sendSvr := &(this.sendSvrList[i])

		sender = sendSvr.senderMail

		if sendSvr.netType == 2 {
			rtnerr = sendSvr.sendByTLS(to, subject, mailtype, body)
		} else {
			rtnerr = sendSvr.sendByNormal(to, subject, mailtype, body)
		}

		if rtnerr == nil {
			return
		}
	}

	return
}

func (this *MailStu) TransToHtml(mailBodyIn []byte) (mailBody []byte, ret int) {

	chars := "UTF-8"
	if this.charType == CharTypeUseGbk {
		chars = "gbk"
	}

	mailBody, ret = TransToHtml(mailBodyIn, chars, "14")

	return
}

//return a smtp client
func (this *sendSvrStu) tlsDial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err

/* to 用 ; 隔开多个收件人
   subject 标题
   mailtype html plain 指明按哪一种格式显示
   body */
func (this *sendSvrStu) sendByTLS(to, subject, mailtype string, body []byte) (err error) {

	chars := "UTF-8"
	if this.charType == CharTypeUseGbk {
		chars = "gbk"
		str1, err1 := CharUtf8ToGBK(subject)
		if err1 == nil {
			subject = str1
		}
		str2, err2 := CharBytesToGBK(body)
		if err2 == nil {
			body = []byte(str2)
		}
	} else if this.charType == CharTypeUseUft8 {
		str1, err1 := CharGbkToUTF8(subject)
		if err1 == nil {
			subject = str1
		}
		str2, err2 := CharBytesToUTF8(body)
		if err2 == nil {
			body = []byte(str2)
		}
	}

	message := ""
	message += fmt.Sprintf("%s: %s\r\n", "From", this.senderMail)
	message += fmt.Sprintf("%s: %s\r\n", "To", to)
	message += fmt.Sprintf("%s: %s\r\n", "Subject", subject)
	message += fmt.Sprintf("%s: %s\r\n", "Content-Type", "text/"+mailtype+"; charset="+chars)
	message += "\r\n"

	message += string(body)

	auth := smtp.PlainAuth("", this._sender_Addr, this.senderPasswd, this.sendSvrAddr)

	//prnLog.Debugf( "%v", StructToJsonStringOne(auth))

	toList := strings.Split(to, ";")

	//create smtp client
	c, err := this.tlsDial(this._sendSvr_AddrPort)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(this._sender_Addr); err != nil {
		return err
	}

	for _, addr := range toList {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}

/* to 用 ; 隔开多个收件人
   subject 标题
   mailtype html plain 指明按哪一种格式显示
   body */
func (this *sendSvrStu) sendByNormal(to, subject, mailtype string, body []byte) error {

	auth := smtp.PlainAuth("", this._sender_Addr, this.senderPasswd, this.sendSvrAddr)

	chars := "UTF-8"
	if this.charType == CharTypeUseGbk {
		chars = "gbk"
		str1, err1 := CharUtf8ToGBK(subject)
		if err1 == nil {
			subject = str1
		}
		str2, err2 := CharBytesToGBK(body)
		if err2 == nil {
			body = []byte(str2)
		}
	} else if this.charType == CharTypeUseUft8 {
		str1, err1 := CharGbkToUTF8(subject)
		if err1 == nil {
			subject = str1
		}
		str2, err2 := CharBytesToUTF8(body)
		if err2 == nil {
			body = []byte(str2)
		}
	}

	// prnLog.Debugf( "%v", StructToJsonStringOne(auth))

	message := ""
	message += fmt.Sprintf("%s: %s\r\n", "From", this.senderMail)
	message += fmt.Sprintf("%s: %s\r\n", "To", to)
	message += fmt.Sprintf("%s: %s\r\n", "Subject", subject)
	message += fmt.Sprintf("%s: %s\r\n", "Content-Type", "text/"+mailtype+"; charset="+chars)
	message += "\r\n"

	message += string(body)

	sendList := strings.Split(to, ";")

	err := smtp.SendMail(this._sendSvr_AddrPort, auth, this._sender_Addr, sendList, []byte(message))

	return err
}

func (this *sendSvrStu) SetCharType(charTypeIn int) {
	this.charType = charTypeIn
	if charTypeIn == CharTypeUseUft8 {
		str1, err1 := CharGbkToUTF8(this.senderMail)
		if err1 == nil {
			this.senderMail = str1
		}
		prnLog.Debugf("senderMail GbkToUTF8, err=%v", err1)
	} else if charTypeIn == CharTypeUseGbk {
		str1, err1 := CharUtf8ToGBK(this.senderMail)
		if err1 == nil {
			this.senderMail = str1
		}
		prnLog.Debugf("senderMail Utf8ToGBK, err=%v", err1)
	}

	return
}

/* charset1 默认 "UTF-8", fontsize 默认 12 */
func TransToHtml(BodyIn []byte, charset1, fontsize string) (BodyOut []byte, ret int) {

	if len(charset1) == 0 {
		charset1 = "UTF-8"
	}
	if len(fontsize) == 0 {
		fontsize = "12"
	}

	bodystr := string(BodyIn)
	bodystr = strings.Replace(bodystr, "\n\r", "<br>", -1)
	bodystr = strings.Replace(bodystr, "\n", "<br>", -1)
	bodystr = strings.Replace(bodystr, " ", "&nbsp;", -1)

	bodystart := `<html>
	<head>
	<meta http-equiv="Content-Type" content="text/html; charset=` + charset1 + `" />
	<title></title>
	</head>
	<body style="font-family:courier;font-size:` + fontsize + `px;color:black;">`
	bodyend := `</body>
</html>`

	BodyOut = []byte(bodystart + bodystr + bodyend)

	return
}

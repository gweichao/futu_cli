package libfunc

import (
	"fmt"
	"testing"
	"time"
)

var (
	smtpSendSvrAddr string = "smtp.163.com"
	smtpSendSvrPort string = "25"

	sendorMail   string = "李四<lisi@163.com>"
	sendorPasswd string = "" // 请赋值真实的密码
)


//go test -v -run="Test_MailByTls"
func _Test_MailByTls(t *testing.T) {

	var Mail1 MailStu
	var Mailsetting []MailsettingStu = []MailsettingStu{
		{SendSvrAddr: "smtp.qq.com", SendSvrPort: "994", NetType: "tls", SenderMail: "李四<lisi@qq.com>", SenderPassword: "***"},
		// {SendSvrAddr: "smtp.163.com", SendSvrPort: "25", NetType: "normal", SenderMail: "李四<lisi@163.com>", SenderPassword: "***"},
	}
	Mail1.AddSendSetting(Mailsetting)

	to := "test@test.com"
	subject := "re [即时]888,2级[S_028]读取本地服务器返回的场景动作失败; 告警数:[rec:13]13.1, 05-09 02:21:55" // "使用Golang发送邮件，MailByTls"

	mailtype := "html" // "plain"

	body := fmt.Sprintf(`<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    </head>
    <body data-lang="zh-cn"> 
        <p>
		<a href="http://hotelmonitor.ibroadlink.com/moncenter/daily_rpt/%v">请点击链接：告警日报 %v</a>
		</p>
	</body>
</html>
`, "daily_rpt_2019-05-09,08.00.html", "2019-05-09,08.00")

	/*fmt.Sprintf(`
	  <html>
	  <body>
	  "%v, %v"
	  </body>
	  </html>
	  `, "Test send to email / 这是一封测试邮件 ", time.Now().Format(formatYMDHMS))*/

	prnLog.Debugf("%v", "send email")

	sender, err := Mail1.Sendmail(to, subject, mailtype, []byte(body)) // html
	if err != nil {
		t.Error(err, "Send mail error! sender="+sender)
	} else {
		prnLog.Debugf("%v,sender=%v", "Send mail success!", sender)
	}
}

//go test -v -run="Test_MailByNormal"
func _Test_MailByNormal(t *testing.T) {

	var Mailsetting []MailsettingStu = []MailsettingStu{
		// {SendSvrAddr: "smtp.qq.com", SendSvrPort: "994", NetType: "tls", SenderMail: "李四<lisi@qq.com>", SenderPassword: "***"},
		{SendSvrAddr: "smtp.163.com", SendSvrPort: "25", NetType: "normal", SenderMail: "李四<lisi@163.com>", SenderPassword: "***"},
	}

	var Mail1 MailStu
	Mail1.AddSendSetting(Mailsetting)

	to := "test@test.com"
	subject := "使用Golang发送邮件，MailByNormal"
	body := fmt.Sprintf(`
                <html>
                <body>
                "%v, %v"
                </body>
                </html>
                `, "Test send to email / 这是一封测试邮件 ", time.Now().Format(formatYMDHMS))

	prnLog.Debugf("%v", "send email")

	sender, err := Mail1.Sendmail(to, subject, "html", []byte(body))
	if err != nil {
		t.Error(err, "Send mail error! sender="+sender)
	} else {
		prnLog.Debugf("%v,sender=%v", "Send mail success!", sender)
	}
}

package email

import (
	"fmt"
	email2 "github.com/Gre-Z/common/email/plugin"
	"net/smtp"
	"net/textproto"
)

const (
	QqExmailHost = "smtp.exmail.qq.com" //QQ企业邮箱域名
	QqExamilAddr = QqExmailHost + ":25" //QQ企业邮箱发送服务器
	QqMailHost   = "smtp.qq.com"        // QQ邮箱域名
	QqMailAddr   = QqMailHost + ":587"  //QQ邮箱发送服务器
	WyMailHost   = "smtp.163.com"
	WyMailAddr   = WyMailHost + ":25"
)

type email struct {
	serverHost string
	sendAddr   string
	sendTo     []string
	author     string
	title      string
	userName   string
	passWord   string
	email      *email2.Email
}

func NewQQExmail(UserName, PassWord string) *email {
	return NewMail(QqExmailHost, QqExamilAddr, UserName, PassWord)
}
func NewQQmail(UserName, PassWord string) *email {
	return NewMail(QqMailHost, QqMailAddr, UserName, PassWord)
}
func New163mail(UserName, PassWord string) *email {
	return NewMail(WyMailHost, WyMailAddr, UserName, PassWord)
}
func NewMail(ServerHost, SendAddr, UserName, PassWord string) *email {
	this := email{
		userName:   UserName,
		passWord:   PassWord,
		serverHost: ServerHost,
		sendAddr:   SendAddr,
	}
	return &this
}
func (this *email) Info(Title, Author string, SendTo []string) *email {
	this.title = Title
	this.author = Author
	this.sendTo = SendTo
	this.email = &email2.Email{
		To:      this.sendTo,
		From:    fmt.Sprintf("%s <%s>", this.author, this.userName),
		Subject: this.title,
		Headers: textproto.MIMEHeader{},
	}
	return this
}

func (this *email) SendHtml(Html string) error {
	this.email.HTML = []byte(Html)
	return this.send()
}

func (this *email) SendText(Text string) error {
	this.email.Text = []byte(Text)
	return this.send()
}

func (this *email) send() error {
	auth := smtp.PlainAuth("", this.userName, this.passWord, this.serverHost)
	return this.email.Send(this.sendAddr, auth)
}

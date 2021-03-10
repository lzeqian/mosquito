package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/jordan-wright/email"
	"gpm/service"
	"gpm/tools"
	"mime"
	"net/smtp"
	"path/filepath"
	"regexp"
	"strings"
)

type EmailController struct {
	beego.Controller
}

func (c *EmailController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	PubInit(c.Controller, ctx, controllerName, actionName, app)
}
func SendToMailOri(host, user, password, fromUser, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + fromUser + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
func SendMail(host string, user string, password string, fromUser string, toUser string, subject string, body []byte, fileName string) error {
	// NewEmail返回一个email结构体的指针
	e := email.NewEmail()
	// 发件人
	e.From = fromUser
	e.Sender = user
	// 收件人(可以有多个)
	e.To = []string{toUser}
	// 邮件主题
	e.Subject = subject
	sendType := checkSendType(fileName)
	//表示普通文本，直接作为html内容发送
	if sendType == 0 {
		e.HTML = body
	}
	//表示转换为html，并且附件方式发送。
	if sendType == 2 {
		e.HTML = body
		htmlBodyByte := tools.ExportToFormat(body, fileName, "html", beego.AppConfig.String("libreofficeTmpPath"))
		bodyBuffer := new(bytes.Buffer)
		bodyBuffer.Write(htmlBodyByte)
		e.HTML = []byte(htmlBodyByte)
		ct := mime.TypeByExtension(filepath.Ext(fileName))
		// 从缓冲中将内容作为附件到邮件中
		e.Attach(bodyBuffer, fileName, ct)
	}
	//表示2进制文件直接附件方式发送。
	if sendType == 1 {
		//Buffer是一个实现了读写方法的可变大小的字节缓冲
		bodyBuffer := new(bytes.Buffer)
		bodyBuffer.Write(body)
		// html形式的消息
		e.HTML = []byte("内容详见附件:(" + fileName + ")")
		ct := mime.TypeByExtension(filepath.Ext(fileName))
		// 从缓冲中将内容作为附件到邮件中
		e.Attach(bodyBuffer, fileName, ct)
	}
	// 发送邮件(如果使用QQ邮箱发送邮件的话，passwd不是邮箱密码而是授权码)
	return e.Send(host, smtp.PlainAuth("", user, password, strings.Split(host, ":")[0]))
}

/**
  根据类型做不通处理
  0 表示普通文本，直接作为html内容发送
  1 表示2进制文件直接附件方式发送。
  2 表示转换为html，并且附件方式发送。
*/
func checkSendType(fileName string) int8 {
	matched, _ := regexp.MatchString("^.*\\.(txt|html|js|css|java|ruby|sh)$", fileName)
	if matched {
		return 0
	}
	//暂时不支持
	//officeMatched,_:=regexp.MatchString("^.*\\.(xls|xlsx|doc|docx|ppt|pptx)$",fileName)
	//if officeMatched{
	//	return 2;
	//}
	return 1
}

/**
  获取子目录结构
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (this *EmailController) SeneMail() {
	data := this.Ctx.Input.RequestBody
	paramData := make(map[string]string)
	json.Unmarshal(data, &paramData)
	token := this.Ctx.Input.Header("Authorization")
	clwas, _ := tools.GetTokenInfo(token)
	subject := paramData["subject"]
	receiver := paramData["receiver"]
	fileDir := paramData["fileDir"]
	fileName := paramData["fileName"]

	userInfo := service.GetUser(clwas.Name)
	sendUser := userInfo["userFullName"].(string) + " <" + userInfo["email"].(string) + ">"
	smtpHost := beego.AppConfig.String("email.smtp.host")
	smtpSender := beego.AppConfig.String("email.smtp.sender")
	smtpPassword := beego.AppConfig.String("email.smtp.password")
	body, err := GetFileSystem(this.Ctx).ReadByte(fileDir, fileName)
	if err != nil {
		ServeJSON(this.Controller, err)
	}
	err = SendMail(smtpHost, smtpSender, smtpPassword, sendUser, receiver, subject, body, fileName)
	if err != nil {
		ServeJSON(this.Controller, err)
	}
	ServeJSON(this.Controller, "")
}

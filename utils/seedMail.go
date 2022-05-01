package utils

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendsCoderMail(subject string, body string,openid string) error {

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "service@ifbiu.com"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{"46982415@qq.com","sevenone_m@163.com","616436112@qq.com"}

	// 设置主题
	em.Subject = subject

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("以下内容为迅捷通用户反馈\n反馈的内容为：\n\n\n"+body+"\n\n\n 反馈者的openid为： \n"+openid)

	//设置服务器相关的配置
	err := em.Send("smtp.ym.163.com:25", smtp.PlainAuth("", "service@ifbiu.com", "wanghao521", "smtp.ym.163.com"))
	if err != nil {
		return err
	}
	return err
}
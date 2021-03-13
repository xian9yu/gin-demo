package utils

//
//import (
//	"admin_system/admin_base/api/common"
//	"admin_system/admin_base/api/internal/svc"
//	"admin_system/admin_base/model/redisPlugin"
//	"context"
//	"fmt"
//	"github.com/tal-tech/go-zero/core/logx"
//	"gopkg.in/gomail.v2"
//	"math/rand"
//	"strconv"
//	"time"
//)

///*
//短信和邮箱验证码的发送及校验
//*/
//type EmailConf struct {
//	EmailAddr       string
//	EmailPort       int
//	EmailBody       string
//	EmailServerHost string
//	EmailSendUser   string
//	EmailPwd        string
//	EmailSubject    string
//}
//
//func MailConf(ec EmailConf) EmailConf {
//	return EmailConf{
//		EmailAddr:       ec.EmailAddr,
//		EmailPort:       ec.EmailPort,
//		EmailBody:       ec.EmailBody,
//		EmailServerHost: ec.EmailServerHost,
//		EmailSendUser:   ec.EmailSendUser,
//		EmailPwd:        ec.EmailPwd,
//		EmailSubject:    ec.EmailSubject,
//	}
//}
//
//// 发送邮件并把验证码存入 redis
//func SendEmail(email string, id int64) error {
//	code := getRand()
//	err := writeMail(code, email)
//	if err == nil {
//		return redisPlugin.RDB.Setex(code, email+strconv.FormatInt(id, 10), 3600) // 过期时间设置为3600s =(1 *time.Hour)
//	}
//	return nil
//}
//
//// 发送邮件
//func writeMail(code, useremail string) error {
//	m := gomail.NewMessage()
//	m.SetAddressHeader("From", EmailAddr, EmailSendUser)     // 发件人
//	m.SetHeader("To", m.FormatAddress(useremail, useremail)) //收件人
//	m.SetHeader("Subject", Subject)                          // 主题
//
//	//body := "您的验证码为：" + code + "  验证码1小时内有效，请尽快完成激活。" // TODO 配置发件body
//	m.SetBody("text/plain", EmailBody)                                    // 正文
//	d := gomail.NewDialer(EmailServerHost, EmailPort, EmailAddr, FromPwd) // 邮件服务器、端口、发件人账号、IMAP/SMTP服务授权密码
//
//	err := d.DialAndSend(m)
//	return err
//}
//
//// 生成 6位随机验证码
//func getRand() string {
//	rand.Seed(time.Now().Unix())
//	randNums := strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) +
//		strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) +
//		strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10))
//	return randNums
//}

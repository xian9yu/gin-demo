package utils

import (
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type SmsInfo struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	SignName        string `json:"sign_name"`      // 短信签名（阿里云后台申请设置）
	TemplateCode    string `json:"template_code"`  // 短信模板代码
	TemplateParam   string `json:"template_param"` // 验证码（golang只需传code，完整短信在阿里云后台设置）
}

// AliSendSMS 阿里云发送短信
func AliSendSMS(sms SmsInfo, phoneNumber, code string) error {
	config := &openapi.Config{
		AccessKeyId:     tea.String(sms.AccessKeyId),     // AccessKey ID
		AccessKeySecret: tea.String(sms.AccessKeySecret), // AccessKey Secret
	}

	config.Endpoint = tea.String("dysmsapi.aliyuncs.com") // 短信服务器
	client, errNewClient := dysmsapi20170525.NewClient(config)
	if errNewClient != nil {
		return errNewClient
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{ // 发送短信参数
		PhoneNumbers:  tea.String(phoneNumber),
		SignName:      tea.String(sms.SignName),
		TemplateCode:  tea.String(sms.TemplateCode),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	resp, errSendSms := client.SendSms(sendSmsRequest)
	if errSendSms != nil {
		return errSendSms
	}

	if *resp.Body.Message != "OK" {
		return errors.New(*resp.Body.Message)
	}
	return nil
}

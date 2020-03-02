package email_service

// ----------------------------------------------------------------------
// 完成SMTP发送
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
// 文档 http://www.voidcn.com/article/p-poayptxe-bwd.html
// ----------------------------------------------------------------------

import (
	"crypto/tls"
	"email_server/pkg/setting"
	"gopkg.in/gomail.v2"
)

type Smtp struct {
	Subject      string   `json:"subject"`
	SenderName   string   `json:"sender_name"`
	Body         string   `json:"body"`
	Receiver     []string `json:"receiver"`
	ReceiverName []string `json:"receiver_name"`
}

func (s *Smtp) Send() error {

	// 内容配置
	m := gomail.NewMessage()
	m.SetAddressHeader("From", setting.SmtpSetting.MAIL_FROM_ADDRESS, s.SenderName)

	no_receiver := false
	if len(s.ReceiverName) == 0 {
		no_receiver = true
	}

	format_emails := []string{}
	for key, _ := range s.Receiver {
		if no_receiver {
			format_emails = append(format_emails, m.FormatAddress(s.Receiver[key], ""))
		} else {
			format_emails = append(format_emails, m.FormatAddress(s.Receiver[key], s.ReceiverName[key]))
		}
	}
	m.SetHeader("To", format_emails...)
	m.SetHeader("Subject", s.Subject)
	m.SetBody("text/html", s.Body)

	// 基础配置
	d := gomail.NewDialer(
		setting.SmtpSetting.MAIL_HOST,
		setting.SmtpSetting.MAIL_PORT,
		setting.SmtpSetting.MAIL_USERNAME,
		setting.SmtpSetting.MAIL_PASSWORD,
	)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: setting.SmtpSetting.MAIL_ENCRYPTION_IS_TLS,
	}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil

}

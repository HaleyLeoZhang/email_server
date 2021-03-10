package bo

import (
	"crypto/tls"
	"github.com/HaleyLeoZhang/email_server/pkg/file"
	"gopkg.in/gomail.v2"
	"strings"
)

type Smtp struct {
	Subject      string   `json:"subject"`
	SenderName   string   `json:"sender_name"`
	Body         string   `json:"body"`
	Receiver     []string `json:"receiver"`
	ReceiverName []string `json:"receiver_name"`
	Attachment   []string `json:"attachment"`
	Remark       []string `json:"remark"`
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

	upload := &Upload{}
	fileList := s.GetAttachmentList()
	for filePath, fileAlias := range fileList {
		if e.UPLOAD_FILE_EXISTS == upload.CheckFile(filePath) {
			m.Attach(filePath, gomail.Rename(fileAlias))
		} else {
			s.Remark = append(s.Remark, "Not Found: "+fileAlias)
		}
	}

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

func (s *Smtp) GetAttachmentList() map[string]string {
	list := make(map[string]string) // file_name file_alias
	upload := &Upload{}

	for _, value := range s.Attachment {
		stringSlice := strings.Split(value, e.UPLOAD_TMP_ALIAS_DELIMITER)
		fileName := stringSlice[0]
		fileAlias := stringSlice[1]
		filePath := upload.GetTmpFilePath(fileName)
		list[filePath] = fileAlias
	}
	return list
}

func (s *Smtp) DeleteAttachmentList() {
	fileList := s.GetAttachmentList()
	upload := &Upload{}

	for filePath, _ := range fileList {
		if e.UPLOAD_FILE_EXISTS == upload.CheckFile(filePath) {
			file.Delete(filePath)
		}
	}
}

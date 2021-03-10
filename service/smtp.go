package service

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
	"github.com/HaleyLeoZhang/email_server/conf"
	"github.com/HaleyLeoZhang/email_server/constant"
	"github.com/HaleyLeoZhang/email_server/model/bo"
	"github.com/HaleyLeoZhang/email_server/pkg/file"
	"gopkg.in/gomail.v2"
	"strings"
)


func (s *Service) SmtpSend(smtp *bo.Smtp) error {

	// 内容配置
	m := gomail.NewMessage()
	m.SetAddressHeader("From",  conf.Conf.Email.Smtp.FromAddr, smtp.SenderName)

	noReceiver := false
	if len(smtp.ReceiverName) == 0 {
		noReceiver = true
	}

	formatEmails := []string{}
	for key, _ := range smtp.Receiver {
		if noReceiver {
			formatEmails = append(formatEmails, m.FormatAddress(smtp.Receiver[key], ""))
		} else {
			formatEmails = append(formatEmails, m.FormatAddress(smtp.Receiver[key], smtp.ReceiverName[key]))
		}
	}
	m.SetHeader("To", formatEmails...)
	m.SetHeader("Subject", smtp.Subject)
	m.SetBody("text/html", smtp.Body)

	fileList := s.SmtpGetAttachmentList(smtp)
	for filePath, fileAlias := range fileList {
		if constant.UPLOAD_FILE_EXISTS == s.UploadCheckFile(filePath) {
			m.Attach(filePath, gomail.Rename(fileAlias))
		} else {
			smtp.Remark = append(smtp.Remark, "Not Found: "+fileAlias)
		}
	}

	// 基础配置
	d := gomail.NewDialer(
		conf.Conf.Email.Smtp.Host,
		conf.Conf.Email.Smtp.Port,
		conf.Conf.Email.Smtp.User,
		conf.Conf.Email.Smtp.Password,
	)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: conf.Conf.Email.Smtp.Tls,
	}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil

}

func (s *Service) SmtpGetAttachmentList(smtp *bo.Smtp) map[string]string {
	list := make(map[string]string) // file_name file_alias
	for _, value := range smtp.Attachment {
		stringSlice := strings.Split(value, constant.UPLOAD_TMP_ALIAS_DELIMITER)
		fileName := stringSlice[0]
		fileAlias := stringSlice[1]
		filePath := s.UploadGetTmpFilePath(fileName)
		list[filePath] = fileAlias
	}
	return list
}

func (s *Service) SmtpDeleteAttachmentList(smtp *bo.Smtp) {
	fileList := s.SmtpGetAttachmentList(smtp)

	for filePath, _ := range fileList {
		if constant.UPLOAD_FILE_EXISTS == s.UploadCheckFile(filePath) {
			file.Delete(filePath)
		}
	}
}

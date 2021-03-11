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
	"github.com/HaleyLeoZhang/email_server/model/vo"
	"github.com/HaleyLeoZhang/email_server/pkg/file"
	"gopkg.in/gomail.v2"
	"strings"
)


func (s *Service) SmtpSend(param *vo.SendEmailRequest) error {
	// 内容配置
	m := gomail.NewMessage()
	m.SetAddressHeader("From",  conf.Conf.Email.Smtp.FromAddr, param.SenderName)

	noReceiver := false
	if len(param.ReceiverName) == 0 {
		noReceiver = true
	}

	lenReceiver := len(param.Receiver)
	formatEmails := make([]string, 0, lenReceiver)
	for key, _ := range param.Receiver {
		if noReceiver {
			formatEmails = append(formatEmails, m.FormatAddress(param.Receiver[key], ""))
		} else {
			formatEmails = append(formatEmails, m.FormatAddress(param.Receiver[key], param.ReceiverName[key]))
		}
	}
	m.SetHeader("To", formatEmails...)
	m.SetHeader("Subject", param.Subject)
	m.SetBody("text/html", param.Body)

	fileList := s.SmtpGetAttachmentList(param.Smtp)
	for filePath, fileAlias := range fileList {
		if constant.UPLOAD_FILE_EXISTS == s.UploadCheckFile(filePath) {
			m.Attach(filePath, gomail.Rename(fileAlias))
		} else {
			param.Remark = append(param.Remark, "Not Found: "+fileAlias)
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

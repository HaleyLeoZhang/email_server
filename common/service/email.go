package service

// ----------------------------------------------------------------------
// 接收外部发送请求
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"context"
	"encoding/json"
	"github.com/HaleyLeoZhang/email_server/common/constant"
	"github.com/HaleyLeoZhang/email_server/common/model/bo"
	"github.com/HaleyLeoZhang/email_server/common/model/po"
	"strings"
)

func (s *Service) DoMessagePush(smtp *bo.Smtp) error {
	payload, err := json.Marshal(smtp)
	if err != nil {
		return err
	}

	err = s.instanceEmailQueue.Push(s, payload)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DoMessagePull() error {
	err := s.instanceEmailQueue.Pull(s, s.doPull)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) doPull(payload []byte) (err error) {
	smtp := &bo.Smtp{}
	ctx := context.Background()

	errJson := json.Unmarshal(payload, smtp)
	if errJson != nil { // 异常数据不存留
		return
	}
	err = s.SmtpSend(smtp)
	// 记录发送邮件的日志
	email := &po.Email{}
	email.Title = smtp.Subject
	email.Content = smtp.Body
	email.SenderName = smtp.SenderName
	email.Receiver = strings.Join(smtp.Receiver, constant.BUSINESS_EMAIL_DELIMITER)
	email.ReceiverName = strings.Join(smtp.ReceiverName, constant.BUSINESS_EMAIL_DELIMITER)
	email.Attachment = strings.Join(smtp.Attachment, constant.UPLOAD_MULIT_FILE)
	email.Remark = strings.Join(smtp.Remark, constant.BUSINESS_EMAIL_DELIMITER)
	email.IsOk = constant.BUSINESS_EMAIL_IS_OK_YES
	if err != nil {
		email.IsOk = constant.BUSINESS_EMAIL_IS_OK_NO
		_ = s.DB.EmailInsert(ctx, nil, email)
		return
	}

	// 删除用过的文件
	s.SmtpDeleteAttachmentList(smtp)

	err = s.DB.EmailInsert(ctx, nil, email)
	if err != nil {
		return
	}

	return
}

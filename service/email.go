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
	"errors"
	"fmt"
	"github.com/HaleyLeoZhang/email_server/constant"
	"github.com/HaleyLeoZhang/email_server/model/bo"
	"github.com/HaleyLeoZhang/email_server/model/po"
	"github.com/HaleyLeoZhang/email_server/util"
	"strings"
)

func (s *Service) DoUpdate(id int, data map[string]interface{}) error {
	whereMap := make(map[string]interface{})
	whereMap["id"] = id
	data["is_ok"] = constant.BUSINESS_EMAIL_IS_OK_YES

	ctx := context.Background()
	_, err := s.DB.EmailUpdate(ctx, nil, whereMap, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DoPushMessage(smtp *bo.Smtp) error {
	payload, err := json.Marshal(smtp)
	if err != nil {
		return err
	}

	q := s.GetEmailQueue()
	err = q.Push(s, payload)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DoPull() error {
	q := s.GetEmailQueue()
	err := q.Pull(s, s.doPull)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) doPull(payload []byte) (err error) {
	smtp := &bo.Smtp{}

	errJson := json.Unmarshal(payload, smtp)
	if errJson != nil { // 异常数据不存留
		return
	}
	err = s.SmtpSend(smtp)
	if err != nil {
		return
	}
	// 记录发送邮件的日志
	email := &po.Email{}
	email.Title = smtp.Subject
	email.Content = smtp.Body
	email.SenderName = smtp.SenderName
	email.Receiver = strings.Join(smtp.Receiver, constant.BUSINESS_EMAIL_DELIMITER)
	email.ReceiverName = strings.Join(smtp.ReceiverName, constant.BUSINESS_EMAIL_DELIMITER)
	email.Attachment = strings.Join(smtp.Attachment, constant.UPLOAD_MULIT_FILE)
	email.Remark = strings.Join(smtp.Remark, constant.BUSINESS_EMAIL_DELIMITER)

	ctx := context.Background()

	// 删除用过的文件
	s.SmtpDeleteAttachmentList(smtp)

	email.IsOk = constant.BUSINESS_EMAIL_IS_OK_YES
	err = s.DB.EmailInsert(ctx, nil, email)
	if err != nil {
		return
	}

	return
}

func (s *Service) getReceiverAndName(receiver string, receiverName string) ([]string, []string, error) {

	receiverArr := strings.Split(receiver, constant.BUSINESS_EMAIL_DELIMITER)

	for _, email := range receiverArr {
		if false == util.CheckEmail(email) {
			return nil, nil, errors.New("receiver 含格式不正确的邮箱地址")
		}
	}

	var receiverNameArr []string
	if "" == receiverName {
		receiverNameArr = []string{}
	} else {
		receiverNameArr = strings.Split(receiverName, constant.BUSINESS_EMAIL_DELIMITER)
		fmt.Printf("%v  %v \n", receiverArr, receiverNameArr)
		if len(receiverArr) != len(receiverNameArr) {
			return nil, nil, errors.New("receiver 与 receiverName 数量不一致")
		}
	}

	return receiverArr, receiverNameArr, nil
}

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
	"github.com/HaleyLeoZhang/email_server/pkg/e"
	"github.com/HaleyLeoZhang/email_server/pkg/queue_engine"
	"github.com/HaleyLeoZhang/email_server/pkg/util"
	"strings"
)

const (
	isOkYes   = 1
	isOkNo    = 0
	delimiter = ","
)

func (s *Service) DoUpdate(id int, data map[string]interface{}) error {
	whereMap := make(map[string]interface{})
	whereMap["id"] = id
	data["is_ok"] = isOkYes

	ctx := context.Background()
	_, err := s.DB.EmailUpdate(ctx, nil, whereMap, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DoPush(data map[string]interface{}) error {
	receiver, receiverName, err := e.getReceiverAndName(
		data["receiver"].(string),
		data["receiverName"].(string),
	)
	if err != nil {
		return err
	}

	smtp := &bo.Smtp{}

	smtp.Subject = data["title"].(string)
	smtp.Body = data["content"].(string)
	smtp.SenderName = data["sender_name"].(string)
	smtp.Receiver = receiver
	smtp.ReceiverName = receiverName
	smtp.Attachment = data["attachment"].([]string)

	payload, err := json.Marshal(smtp)
	if err != nil {
		return err
	}

	q := queue_engine.GetEmailQueue()
	err = q.Push(s, payload)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DoPull() error {

	q := queue_engine.GetEmailQueue()
	err := q.Pull(s, s.doPull)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) doPull(payload []byte) error {
	smtp := &bo.Smtp{}

	err := json.Unmarshal(payload, smtp)
	if err != nil {
		return nil
	}
	err = s.SmtpSend(smtp)

	email := &po.Email{}
	email.Title = smtp.Subject
	email.Content = smtp.Body
	email.SenderName = smtp.SenderName
	email.Receiver = strings.Join(smtp.Receiver, ",")
	email.ReceiverName = strings.Join(smtp.ReceiverName, ",")
	email.Attachment = strings.Join(smtp.Attachment, constant.UPLOAD_MULIT_FILE)
	email.Remark = strings.Join(smtp.Remark, ",")

	ctx := context.Background()
	if err != nil {
		email.IsOk = isOkNo
		_ = s.DB.EmailInsert(ctx, nil, email)
		return err
	}

	// 删除用过的文件
	s.SmtpDeleteAttachmentList(smtp)

	email.IsOk = isOkYes
	err = s.DB.EmailInsert(ctx, nil, email)
	if err != nil {
		return err
	}

	// fmt.Printf("发送邮件成功: %v \n", payload)

	return nil
}

func (s *Service) getReceiverAndName(receiver string, receiverName string) ([]string, []string, error) {

	receiverArr := strings.Split(receiver, delimiter)

	for _, email := range receiverArr {
		if false == util.CheckEmail(email) {
			return nil, nil, errors.New("receiver 含格式不正确的邮箱地址")
		}
	}

	var receiverNameArr []string
	if "" == receiverName {
		receiverNameArr = []string{}
	} else {
		receiverNameArr = strings.Split(receiverName, delimiter)
		fmt.Printf("%v  %v \n", receiverArr, receiverNameArr)
		if len(receiverArr) != len(receiverNameArr) {
			return nil, nil, errors.New("receiver 与 receiverName 数量不一致")
		}
	}

	return receiverArr, receiverNameArr, nil
}

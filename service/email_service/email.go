package email_service

// ----------------------------------------------------------------------
// 接收外部发送请求
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"email_server/models"
	"email_server/pkg/queue"
    "email_server/pkg/util"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Email struct{}

const (
	isOkYes   = 1
	isOkNo    = 0
	delimiter = ","
)

func (e *Email) DoUpdate(id int, data map[string]interface{}) error {
	email := new(models.Email)
	email.ID = id

	data["is_ok"] = isOkYes

	err := email.Update(data)
	if err != nil {
		return err
	}
	return nil
}

func (e *Email) DoPush(data map[string]interface{}) error {
	receiver, receiver_name, err := e.getReceiverAndName(
		data["receiver"].(string),
		data["receiver_name"].(string),
	)
	if err != nil {
		return err
	}

	smtp := &Smtp{}

	smtp.Subject = data["title"].(string)
	smtp.Body = data["content"].(string)
	smtp.SenderName = data["sender_name"].(string)
	smtp.Receiver = receiver
	smtp.ReceiverName = receiver_name

	payload, err := json.Marshal(smtp)
	if err != nil {
		return err
	}

	q := queue.GetEmailQueue()
	q.Payload = string(payload)
	err = q.Push()
	if err != nil {
		return err
	}
	return nil
}

func (e *Email) DoPull() error {

	q := queue.GetEmailQueue()
	err := q.Pull(doPull)
	if err != nil {
		return err
	}
	return nil
}

func doPull(payload string) error {
	smtp := &Smtp{}

	err := json.Unmarshal([]byte(payload), smtp)
	if err != nil {
		// fmt.Printf("消费失败,格式不正确,丢弃当前消息: %v\n", payload)
		return nil
	}
	fmt.Printf("正在发送邮件: %v \n", payload)
	err = smtp.Send()

	email := new(models.Email)
	email.Title = smtp.Subject
	email.Content = smtp.Body
	email.SenderName = smtp.SenderName
	email.Receiver = strings.Join(smtp.Receiver, ",")
	email.ReceiverName = strings.Join(smtp.ReceiverName, ",")
	// email.Attachment = data["attachment"]

	if err != nil {
		email.IsOk = isOkNo
		_ = email.Create()
		fmt.Printf("发送失败: %v \n", payload)
		return err
	}

	email.IsOk = isOkYes
	err = email.Create()
	if err != nil {
		return err
	}

	fmt.Printf("发送邮件成功: %v \n", payload)

	return nil
}

func (e *Email) getReceiverAndName(receiver string, receiver_name string) ([]string, []string, error) {

	receiver_arr := strings.Split(receiver, delimiter)

    for _, email := range receiver_arr {
        if false == util.RegEmail(email) {
            return nil, nil, errors.New("receiver 含格式不正确的邮箱地址") 
        }
    }

	var receiver_name_arr []string
	if "" == receiver_name {
		receiver_name_arr = []string{}
	} else {
		receiver_name_arr = strings.Split(receiver_name, delimiter)
		fmt.Printf("%v  %v \n", receiver_arr, receiver_name_arr)
		if len(receiver_arr) != len(receiver_name_arr) {
			return nil, nil, errors.New("receiver 与 receiver_name 数量不一致")
		}
	}

	return receiver_arr, receiver_name_arr, nil
}

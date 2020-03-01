package email_service

// ----------------------------------------------------------------------
// 接收外部发送请求
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"email_server/models"
)

type Email struct{}

const (
	isOkYes = 1
	isOkNo  = 0
)

func (e *Email) DoCreate(data map[string]interface{}) error {
	email := new(models.Email)
	email.Title = data["title"].(string)
	email.Content = data["content"].(string)
	email.SenderName = data["sender_name"].(string)
	email.Receiver = data["receiver"].(string)
	email.ReceiverName = data["receiver_name"].(string)
	email.Attachment = data["attachment"].(string)
	email.IsOk = isOkNo

	err := email.Create()
	if err != nil {
		return err
	}
	return nil
}

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

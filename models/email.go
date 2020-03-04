package models

// ----------------------------------------------------------------------
// 邮件服务-模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
import (
	"email_server/pkg/e"
)

type Email struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	SenderName   string `json:"sender_name"`
	Receiver     string `json:"receiver"`
	ReceiverName string `json:"receiver_name"`
	Attachment   string `json:"attachment"`
	Remark       string `json:"remark"`
	IsOk         int    `json:"is_ok"`
}

func (email *Email) Create() error {
	if err := db.Create(&email).Error; err != nil {
		return err
	}
	return nil
}

func (email *Email) Update(data map[string]interface{}) error {

	err := db.Model(email).Where("id = ? AND is_deleted = ? ", email.ID, e.DATA_IS_DELETED_NO).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}

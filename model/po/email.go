package po

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

//数据表---必需
func (e *Email) TableName() string {
	return "email"
}

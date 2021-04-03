package bo

type Smtp struct {
	Subject      string   `json:"subject"`
	SenderName   string   `json:"sender_name"`
	Body         string   `json:"body"`
	Receiver     []string `json:"receiver"`
	ReceiverName []string `json:"receiver_name"`
	Attachment   []string `json:"attachment"`
	Remark       []string `json:"remark"`
}

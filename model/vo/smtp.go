package vo

import (
	"github.com/HaleyLeoZhang/email_server/constant"
	"github.com/HaleyLeoZhang/email_server/model/bo"
	"github.com/HaleyLeoZhang/email_server/util"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"strings"
)

type SendEmailRequest struct {
	Title        string `form:"title"  binding:"gte=1,lte=255"`
	Content      string `form:"content"  binding:"gte=1,lte=65535"`
	Receiver     string `form:"receiver"  binding:"gte=1,lte=2000"`
	ReceiverName string `form:"receiver_name" binding:"gte=1,lte=1000"`
	SenderName   string `form:"sender_name" binding:"gte=1,lte=50"`
}

func (s *SendEmailRequest) GetReceiverAndName() (receiverList []string, receiverNameList []string, err error) {
	receiverList = strings.Split(s.Receiver, constant.BUSINESS_EMAIL_DELIMITER)
	receiverNameList = make([]string, 0)

	lenReceiverList := len(receiverList)
	if lenReceiverList == 0 {
		err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_PARAM_INVALID, Message: "receiver 列表为空"}
		return
	}

	for _, email := range receiverList {
		if false == util.CheckEmail(email) {
			err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_PARAM_INVALID, Message: "receiver 含格式不正确的邮箱地址"}
			return
		}
	}

	if "" == s.ReceiverName {
		return
	}
	receiverNameList = strings.Split(s.ReceiverName, constant.BUSINESS_EMAIL_DELIMITER)
	if lenReceiverList != len(receiverNameList) {
		xlog.Warnf("GetReceiverAndName Warn receiverList(%v) receiverNameList(%v)", receiverList, receiverNameList)
		err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_PARAM_INVALID, Message: "receiver 与 receiverName 数量不一致"}
		return
	}
	return
}

func (s *SendEmailRequest) GetSmtpObject() (smtp *bo.Smtp, err error) {
	smtp = &bo.Smtp{}
	err = nil

	smtp.Subject = s.Title
	smtp.Body = s.Content
	smtp.SenderName = s.SenderName

	smtp.Receiver, smtp.ReceiverName, err = s.GetReceiverAndName()
	if err != nil {
		return
	}
	return
}

package vo

import (
	"github.com/HaleyLeoZhang/email_server/constant"
	"github.com/HaleyLeoZhang/email_server/model/bo"
	"github.com/HaleyLeoZhang/email_server/pkg/util"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"strings"
)

type SendEmailRequest struct {
	*bo.Smtp
}

// 检测邮件基础信息
func (s *SendEmailRequest) ValidateBaseInfo() (err error){
	valid.MinSize(data["title"], 1, "title")
	valid.MaxSize(data["title"], 255, "title")
	valid.MinSize(data["content"], 1, "content")
	valid.MaxSize(data["content"], 65535, "content")
	valid.MinSize(data["sender_name"], 1, "sender_name")
	valid.MaxSize(data["sender_name"], 50, "sender_name")
	valid.MinSize(data["receiver"], 1, "receiver")
	valid.MaxSize(data["receiver"], 2000, "receiver")
	valid.MaxSize(data["receiver_name"], 1000, "receiver_name")
}

func (s *SendEmailRequest) CheckReceiverAndName(receiver string, receiverName string) (err error) {
	receiverList := strings.Split(receiver, constant.BUSINESS_EMAIL_DELIMITER)
	receiverNameList := make([]string, 0)

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

	if "" == receiverName {
		return
	}
	receiverNameList = strings.Split(receiverName, constant.BUSINESS_EMAIL_DELIMITER)
	if lenReceiverList != len(receiverNameList) {
		xlog.Warnf("GetReceiverAndName Warn receiverList(%v) receiverNameList(%v)", receiverList, receiverNameList)
		err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_PARAM_INVALID, Message: "receiver 与 receiverName 数量不一致"}
		return
	}
	// 写入参数
	s.Receiver = receiverList
	s.ReceiverName = receiverNameList
	return
}

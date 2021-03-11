package http

// ----------------------------------------------------------------------
// 邮件服务
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"fmt"
	"github.com/HaleyLeoZhang/email_server/constant"
	"github.com/HaleyLeoZhang/email_server/model/vo"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

/**
 * @api {post} /api/email/send 发送邮件
 * @apiName send
 * @apiGroup Email
 *
 * @apiParam {string} title 邮件标题
 * @apiParam {string} content 邮件内容
 * @apiParam {string} sender_name 发件人昵称
 * @apiParam {string} receiver 收件人邮箱,多个以英文逗号隔开
 * @apiParam {string} receiver_name 收件人昵称,可不填,多个以英文逗号隔开
 * @apiParam {file}   attachment[] 多个附件请使用相同变量名.请使用 form-data 进行传输
 *
 * @apiHeaderExample {json} Header-Example:
 * {
 *     "Content-Type": "multipart/form-data"
 * }
 *
 * @apiDescription  发送邮件
 *
 * @apiVersion 1.0.0
 * @apiSuccessExample Success-Response:
 * {
 *     "code": 200,
 *     "message": "success",
 *     "data": null
 * }
 * @apiErrorExample Error-Response:
 * {
 *     "code": 400,
 *     "message": "Param is invalid",
 *     "data": null
 * }
 */
func EmailSend(c *gin.Context) {
	xGin := xgin.NewGin(c)

	// - 获取基础信息
	param := &vo.SendEmailRequest{}
	err := c.Bind(param)
	if err != nil {
		err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_PARAM_INVALID, Message: "Param is invalid"}
		xGin.Response(err, nil)
		return
	}
	smtp, err := param.GetSmtpObject()
	if err != nil {
		xGin.Response(err, nil)
		return
	}
	// - 获取上传的文件
	form, err := c.MultipartForm()
	if err != nil {
		xGin.Response(err, nil)
		return
	}
	files := form.File["attachment[]"]
	smtp.Attachment, err = emailGetFiles(c, files)

	err = srv.DoMessagePush(smtp)

	if err != nil {
		xGin.Response(err, nil)
		return
	}
	xGin.Response(err, nil)
	return
}

func emailGetFiles(c *gin.Context, files []*multipart.FileHeader) (attachment []string, err error) {
	for _, file := range files {
		fileTmpPath, fileTmpName := srv.UploadCreateTmpFile()
		fileAlias := file.Filename
		err = c.SaveUploadedFile(file, fileTmpPath)
		if err != nil {
			return
		}
		oneTmpInfo := fmt.Sprintf("%s%s%s", fileTmpName, constant.UPLOAD_TMP_ALIAS_DELIMITER, fileAlias)
		attachment = append(attachment, oneTmpInfo)
	}
	return
}

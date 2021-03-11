package http

// ----------------------------------------------------------------------
// 邮件服务
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"github.com/HaleyLeoZhang/email_server/constant"
	"github.com/HaleyLeoZhang/email_server/model/vo"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"fmt"
	"github.com/HaleyLeoZhang/email_server/pkg/app"
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
 *     "code": 1001,
 *     "message": "请求参数错误",
 *     "data": [
 *         "receiver 含格式不正确的邮箱地址"
 *     ]
 * }
 */
func EmailSend(c *gin.Context) {
	xGin := xgin.NewGin(c)

	data := make(map[string]interface{})
	data["title"] = com.StrTo(c.PostForm("title")).String()
	data["content"] = com.StrTo(c.PostForm("content")).String()
	data["sender_name"] = com.StrTo(c.PostForm("sender_name")).String()
	data["receiver"] = com.StrTo(c.PostForm("receiver")).String()
	data["receiver_name"] = com.StrTo(c.PostForm("receiver_name")).String()

	// 获取基础信息
	param := &vo.ComicListParam{}
	err := c.Bind(param)
	if err != nil {
		err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_PARAM_INVALID, Message: "Param is invalid"}
		xGin.Response(err, nil)
		return
	}
	// Multipart form --- TODO 2020-3-4 23:51:21
	form, err := c.MultipartForm()
	if err != nil {
		xGin.Response(err, nil)
		return
	}
	files := form.File["attachment[]"]
	attachment := make([]string, 0)
	for _, file := range files {
		fileTmpPath, fileTmpName := srv.UploadCreateTmpFile()
		fileAlias := file.Filename
		if err := c.SaveUploadedFile(file, fileTmpPath); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		oneTmpInfo := fmt.Sprintf("%s%s%s", fileTmpName, constant.UPLOAD_TMP_ALIAS_DELIMITER, fileAlias)
		attachment = append(attachment, oneTmpInfo)
	}
	data["attachment"] = attachment

	valid := validation.Validation{}
	valid.MinSize(data["title"], 1, "title")
	valid.MaxSize(data["title"], 255, "title")
	valid.MinSize(data["content"], 1, "content")
	valid.MaxSize(data["content"], 65535, "content")
	valid.MinSize(data["sender_name"], 1, "sender_name")
	valid.MaxSize(data["sender_name"], 50, "sender_name")
	valid.MinSize(data["receiver"], 1, "receiver")
	valid.MaxSize(data["receiver"], 2000, "receiver")
	valid.MaxSize(data["receiver_name"], 1000, "receiver_name")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, constant.INVALID_PARAMS, nil)
		return
	}
	param := &vo.SendEmailRequest{}
	err = param.CheckReceiverAndName()

	err = srv.DoPush(data)

	if err != nil {
		errInfo := []string{err.Error()}
		appG.Response(http.StatusInternalServerError, constant.INVALID_PARAMS, errInfo)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

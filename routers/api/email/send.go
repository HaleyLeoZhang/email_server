package email

// ----------------------------------------------------------------------
// 邮件服务
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"email_server/pkg/app"
	"email_server/pkg/e"
	"email_server/service/email_service"
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
func Send(c *gin.Context) {
	appG := app.Gin{C: c}

	data := make(map[string]interface{})
	data["title"] = com.StrTo(c.PostForm("title")).String()
	data["content"] = com.StrTo(c.PostForm("content")).String()
	data["sender_name"] = com.StrTo(c.PostForm("sender_name")).String()
	data["receiver"] = com.StrTo(c.PostForm("receiver")).String()
	data["receiver_name"] = com.StrTo(c.PostForm("receiver_name")).String()

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["attachment"]

	for _, file := range files {
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}
	// data["attachment"] = ""

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
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	service := email_service.Email{}
	err := service.DoPush(data)

	if err != nil {
		err_info := []string{err.Error()}
		appG.Response(http.StatusInternalServerError, e.INVALID_PARAMS, err_info)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

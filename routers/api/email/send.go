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
	// "email_server/pkg/setting"
	// "email_server/pkg/util"
	"email_server/service/email_service"
	// "email_server/pkg/logging"
	// "fmt"
)

/**
 * @api {post} /api/email/send 发送邮件
 * @apiName send
 * @apiGroup Email
 *
 * @apiParam {int} page 页码
 *
 * @apiDescription  发送邮件
 *
 * @apiVersion 1.0.0
 * @apiSuccessExample Success-Response:
 */
func Send(c *gin.Context) {
	appG := app.Gin{C: c}

	data := make(map[string]interface{})
	data["title"] = com.StrTo(c.PostForm("title")).String()
	data["content"] = com.StrTo(c.PostForm("content")).String()
	data["sender_name"] = com.StrTo(c.PostForm("sender_name")).String()
	data["receiver"] = com.StrTo(c.PostForm("receiver")).String()
	data["receiver_name"] = com.StrTo(c.PostForm("receiver_name")).String()

	// // 获取解析后表单
	// form, _ := c.MultipartForm()
	// //这里是多文件上传 在之前单文件upload上传的基础上加 [] 变成upload[] 类似文件数组的意思
	// files := form.File["files[]"]
	// //循环存文件到服务器本地
	// for _, file := range files {
	// 	c.SaveUploadedFile(file, file.Filename)
	// }
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

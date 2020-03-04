package email_service

// ----------------------------------------------------------------------
// 完成SMTP发送
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
// 文档 http://www.voidcn.com/article/p-poayptxe-bwd.html
// ----------------------------------------------------------------------

import (
	"email_server/pkg/e"
	"email_server/pkg/file"
	"email_server/pkg/setting"
	"email_server/pkg/util"
	"fmt"
)

type Upload struct{}

func (u *Upload) GetUploadDir() string {
	dir := fmt.Sprintf("%s%s/upload", setting.AppSetting.RuntimeRootPath, setting.AppSetting.UPLOAD_DIR)
	err := file.IsNotExistMkDir(dir)
	if err != nil {
		panic(fmt.Sprintf("文件夹创建失败: %s", err.Error()))
	}
	return dir
}

func (u *Upload) GetTmpFilePath(name string) string {
	dir := u.GetUploadDir()
	return fmt.Sprintf("%s/%s", dir, name)
}

func (u *Upload) CreateTmpFile() (string, string) {
	name := util.GetUuid()
	return u.GetTmpFilePath(name), name
}

func (u *Upload) DeleteTmpFile(name string) {
	filePath := u.GetTmpFilePath(name)
	file.Delete(filePath)
}

func (u *Upload) CheckFile(filePath string) bool {
	if true == file.CheckNotExist(filePath) {
		return e.UPLOAD_FILE_NOT_FOUND
	}
	return e.UPLOAD_FILE_EXISTS
}

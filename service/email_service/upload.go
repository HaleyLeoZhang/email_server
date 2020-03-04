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
    "email_server/pkg/setting"
)

type Upload struct {}

func (u *Upload) GetUploadDir() error {
    return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.UPLOAD_DIR)
}

func (u *Upload) GetUploadPathFile() error {
    path := u.GetUploadDir()
    return fmt.Sprintf("%s/file-%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.UPLOAD_DIR)
}


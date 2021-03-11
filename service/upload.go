package service

// ----------------------------------------------------------------------
// 完成SMTP发送
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
// 文档 http://www.voidcn.com/article/p-poayptxe-bwd.html
// ----------------------------------------------------------------------

import (
	"fmt"
	"github.com/HaleyLeoZhang/email_server/conf"
	"github.com/HaleyLeoZhang/email_server/constant"
	"github.com/HaleyLeoZhang/email_server/util"
)

func (s *Service) UploadGetUploadDir() string {
	dir := conf.Conf.Email.UploadFile.Dir
	err := util.IsNotExistMkDir(dir)
	if err != nil {
		panic(fmt.Sprintf("文件夹创建失败: %s", err.Error()))
	}
	return dir
}

func (s *Service) UploadGetTmpFilePath(name string) string {
	dir := s.UploadGetUploadDir()
	return fmt.Sprintf("%s/%s", dir, name)
}

func (s *Service) UploadCreateTmpFile() (string, string) {
	name := util.GetUuid()
	return s.UploadGetTmpFilePath(name), name
}

func (s *Service) UploadDeleteTmpFile(name string) {
	filePath := s.UploadGetTmpFilePath(name)
	util.Delete(filePath)
}

func (s *Service) UploadCheckFile(filePath string) bool {
	if true == util.CheckNotExist(filePath) {
		return constant.UPLOAD_FILE_NOT_FOUND
	}
	return constant.UPLOAD_FILE_EXISTS
}

package util

import (
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/pkg/errors"

	// "io/ioutil"
	// "mime/multipart"
	"os"
	// "path"
)

// CheckNotExist check if the file exists
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission check if the file has permission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir create a directory
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	return nil
}

// Delete One file
func Delete(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		err = errors.WithStack(err)
		xlog.Errorf("err(%+v)", err)
	}
}

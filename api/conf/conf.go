package conf

import (
	"flag"
	"github.com/HaleyLeoZhang/email_server/common/model/bo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var (
	Conf     = &bo.Config{}
	confPath string
)

func init() {
	flag.StringVar(&confPath, "conf", "", "conf values")
}

func Init() (err error) {
	var yamlFile string
	if confPath != "" {
		yamlFile, err = filepath.Abs(confPath)
	} else {
		yamlFile, err = filepath.Abs("./api/build/app.yaml")
	}
	if err != nil {
		return
	}
	yamlRead, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlRead, Conf)
	if err != nil {
		return
	}
	return
}

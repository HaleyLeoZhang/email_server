package conf

import (
	"flag"
	"github.com/HaleyLeoZhang/email_server/model/bo"
	"github.com/HaleyLeoZhang/go-component/driver/xrabbitmq"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"

	"github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/go-component/driver/httpserver"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/driver/xredis"
)

var (
	Conf     = &Config{}
	confPath string
)

// Config struct
type Config struct {
	ServiceName string             `yaml:"serviceName" json:"serviceName"`
	HttpServer  *httpserver.Config `yaml:"httpServer" json:"httpServer"`
	Gin         *xgin.Config       `yaml:"gin" json:"gin"`
	DB          *db.Config         `yaml:"db" json:"db"`
	Redis       *xredis.Config     `yaml:"redis" json:"redis"`
	Log         *xlog.Config       `yaml:"log" json:"log"`
	RabbitMq    *xrabbitmq.Config  `yaml:"rabbit_mq" json:"rabbit_mq"`
	Email       *bo.ConfigEmail    `yaml:"email" json:"email"`
}

func init() {
	flag.StringVar(&confPath, "conf", "", "conf values")
}

func Init() (err error) {
	var yamlFile string
	if confPath != "" {
		yamlFile, err = filepath.Abs(confPath)
	} else {
		yamlFile, err = filepath.Abs("../build/app.yaml")
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
	go load()
	return
}

func load() {
	// 动态加载配置
}

package conf

import (
	"flag"
	"github.com/HaleyLeoZhang/email_server/model/bo"
	"github.com/HaleyLeoZhang/go-component/driver/xkafka"
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
	RabbitMq    *xrabbitmq.Config  `yaml:"rabbitMq" json:"rabbitMq"`
	Kafka       *xkafka.Config     `yaml:"kafka" json:"kafka"`
	KafkaTopic  *KafkaTopicOne     `yaml:"kafkaTopic" json:"kafkaTopic"` // 如果使用kafka,请配置Topic名称到这个字段
	Email       *bo.ConfigEmail    `yaml:"email" json:"email"`
}

type KafkaTopicOne struct {
	Group     string `yaml:"group"`     // 消费者组名称
	TopicList []string `yaml:"topicList"` // Topic 列表
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

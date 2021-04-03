package bo

import (
	"github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/go-component/driver/httpserver"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/HaleyLeoZhang/go-component/driver/xkafka"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/driver/xrabbitmq"
	"github.com/HaleyLeoZhang/go-component/driver/xredis"
)

type Config struct {
	ServiceName string               `yaml:"serviceName"`
	HttpServer  *httpserver.Config   `yaml:"httpServer"`
	Gin         *xgin.Config         `yaml:"gin"`
	DB          *db.Config           `yaml:"db"`
	Redis       *xredis.Config       `yaml:"redis"`
	Log         *xlog.Config         `yaml:"log"`
	RabbitMq    *xrabbitmq.Config    `yaml:"rabbitMq"`
	Kafka       *xkafka.Config       `yaml:"kafka"`
	KafkaTopic  *ConfigKafkaTopicOne `yaml:"kafkaTopic"` // 如果使用kafka,请配置Topic名称到这个字段
	Email       *ConfigEmail         `yaml:"email"`
}

type ConfigEmail struct {
	Smtp        ConfigSmtp       `yaml:"smtp"`
	Driver      string           `yaml:"driver"`
	Consumer    int              `yaml:"consumer"`
	BatchNumber int              `yaml:"batchNumber"`
	UploadFile  ConfigUploadFile `yaml:"uploadFile"`
}

type ConfigSmtp struct {
	Port     int    `yaml:"port"`
	Tls      bool   `yaml:"tls"`
	FromAddr string `yaml:"fromAddr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}

type ConfigUploadFile struct {
	Dir string `yaml:"dir"`
}
type ConfigKafkaTopicOne struct {
	Group     string   `yaml:"group"`     // 消费者组名称
	TopicList []string `yaml:"topicList"` // Topic 列表
}

package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
	// "os"
	// "fmt"
)

type App struct {
	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Smtp struct {
	MAIL_HOST              string
	MAIL_FROM_ADDRESS      string
	MAIL_USERNAME          string
	MAIL_PASSWORD          string
	MAIL_PORT              int
	MAIL_ENCRYPTION_IS_TLS bool
}

var SmtpSetting = &Smtp{}

type AMQP struct {
	HOST                 string
	PORT                 int
	USER                 string
	PASSWORD             string
	VHOST                string
	CHANNEL_NUMBER       int
}

var AMQPSetting = &AMQP{}

type Queue struct {
	DRIVER string
}

var QueueSetting = &Queue{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var (
		err           error
		ini_file_path string
	)

	// if len(os.Args) > 1 {
	//     ini_file_path = os.Args[1] // args 第一个片 是文件路径
	// }else{
	//     ini_file_path = "/usr/local/etc/mail.ops.hlzblog.top.ini" // 默认读取
	// }
	ini_file_path = "/usr/local/etc/mail.ops.hlzblog.top.ini" // 默认读取

	cfg, err = ini.Load(ini_file_path)

	if err != nil {
		log.Fatalf("setting.Setup, fail to parse '%v': %v", ini_file_path, err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("smtp", SmtpSetting)
	mapTo("amqp", AMQPSetting)
	mapTo("queue", QueueSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

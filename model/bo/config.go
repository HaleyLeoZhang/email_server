package bo

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

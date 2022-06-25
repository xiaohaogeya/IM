package conf

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName      string         `json:"app_name"`
	AppModel     string         `json:"app_model"`
	AppHost      string         `json:"app_host"`
	AppPort      string         `json:"app_port"`
	SecretKey    string         `json:"secret_key"`
	ReadTimeout  int            `json:"read_timeout"`
	WriteTimeout int            `json:"write_timeout"`
	Database     DatabaseConfig `json:"database"`
	RedisConfig  RedisConfig    `json:"redis_config"`
	JWT          JWT            `json:"jwt"`
}

// DatabaseConfig 数据库属性定义
type DatabaseConfig struct {
	Driver            string `json:"driver"`
	User              string `json:"user"`
	Password          string `json:"password"`
	Host              string `json:"host"`
	Port              string `json:"port"`
	DbName            string `json:"db_name"`
	ChartSet          string `json:"charset"`
	ShowSql           bool   `json:"show_sql"`
	SetMaxIdleConnNum int    `json:"set_max_idle_conn_num"`
	SetMaxOpenConnNum int    `json:"set_max_open_conn_num"`
}

// RedisConfig Redis属性定义
type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

// JWT  属性定义
type JWT struct {
	ExpireAt int    `json:"expire_at"`
	Issuer   string `json:"issuer"`
}

var AppConfig *Config

// InitConfig 初始化配置
func InitConfig() (err error) {
	file, err := os.Open("./conf/conf.json")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&AppConfig)
	return
}

func init() {
	err := InitConfig()
	if err != nil {
		panic("初始化项目配置失败")
	}
}

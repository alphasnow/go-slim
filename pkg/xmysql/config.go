package xmysql

import (
	"fmt"
	"gorm.io/gorm/logger"
	"time"
)

type Config struct {
	Base   BaseConfig  `yaml:",inline"`
	Write  DSNConfig   `yaml:",inline"`
	Reads  []DSNConfig `yaml:"reads"`
	Logger LogConfig   `yaml:",inline"`
}

func NewConfig() *Config {
	return &Config{
		Base: BaseConfig{
			Debug:         false,
			TablePrefix:   "",
			SingularTable: true,
			MaxLifetime:   1800,
			MaxIdleConns:  20,
			MaxOpenConns:  60,
		},
		Write: DSNConfig{
			Host:       "127.0.0.1",
			Port:       3306,
			Database:   "tests",
			Username:   "root",
			Password:   "",
			Parameters: "charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true",
		},
		Reads: []DSNConfig{},
		Logger: LogConfig{
			SlowThreshold: 3,
			LogLevel:      "warn",
		},
	}
}

type BaseConfig struct {
	Debug                bool   `yaml:"debug" env:"MYSQL_DEBUG"`
	TablePrefix          string `yaml:"tablePrefix" env:"MYSQL_TABLE_PREFIX"`
	SingularTable        bool   `yaml:"singularTable"`
	MaxLifetime          int    `yaml:"maxLifetime"`
	MaxOpenConns         int    `yaml:"maxOpenConns"`
	MaxIdleConns         int    `yaml:"maxIdleConns"`
	EnableAutoMigrate    bool   `yaml:"enableAutoMigrate"`
	EnableCreateDatabase bool   `yaml:"enableCreateDatabase" env:"MYSQL_ENABLE_CREATE_DATABASE"`
}
type DSNConfig struct {
	Host       string `yaml:"host" env:"MYSQL_HOST"`
	Port       int    `yaml:"port" env:"MYSQL_PORT"`
	Database   string `yaml:"database" env:"MYSQL_DATABASE"`
	Username   string `yaml:"username" env:"MYSQL_USERNAME"`
	Password   string `yaml:"password" env:"MYSQL_PASSWORD"`
	Parameters string `yaml:"parameters" env:"MYSQL_PARAMETERS"`
}

func (c *DSNConfig) DNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Database, c.Parameters)
}
func (c *DSNConfig) DB() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/?%s", c.Username, c.Password, c.Host, c.Port, c.Parameters)
}

type LogConfig struct {
	SlowThreshold int    `yaml:"slowThreshold" env:"MYSQL_LOG_SLOW"`
	LogLevel      string `yaml:"logLevel" env:"MYSQL_LOG_LEVEL"`
}

func (c LogConfig) GetSlowThreshold() time.Duration {
	return time.Duration(c.SlowThreshold) * time.Second
}
func (c LogConfig) GetLogLevel() logger.LogLevel {
	switch c.LogLevel {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Warn
	default:
		return logger.Error
	}
}

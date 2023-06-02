package xredis

import "fmt"

type Config struct {
	Host   string `yaml:"host" env:"REDIS_HOST"`
	Port   int    `yaml:"port" env:"REDIS_PORT"`
	Pass   string `yaml:"pass" env:"REDIS_PASS"`
	DB     int    `yaml:"db" env:"REDIS_DB"`
	Prefix string `yaml:"prefix" env:"REDIS_PREFIX"`
}

func NewConfig() *Config {
	return &Config{
		Host:   "127.0.0.1",
		Port:   6379,
		Pass:   "",
		DB:     0,
		Prefix: "",
	}
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

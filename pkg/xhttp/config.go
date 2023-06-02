package xhttp

import "fmt"

type Config struct {
	Debug        bool   `yaml:"debug" env:"HTTP_DEBUG"`
	Host         string `yaml:"host" env:"HTTP_HOST"`
	Port         int    `yaml:"port" env:"HTTP_PORT"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
	IdleTimeout  int    `yaml:"idleTimeout"`
	UseNginx     bool   `yaml:"useNginx" env:"HTTP_USE_NGINX"`
	UseTLS       bool   `yaml:"useTLS" env:"HTTP_USE_TLS"`
	MaxPostSize  int    `yaml:"maxPostSize"`
}

func NewConfig() *Config {
	return &Config{
		Debug:        false,
		Host:         "127.0.0.1",
		Port:         8080,
		ReadTimeout:  10,
		WriteTimeout: 30,
		IdleTimeout:  30,
		UseNginx:     true,
		MaxPostSize:  10 * 1024 * 1024,
	}
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
func (c *Config) GetUrl() string {
	host := c.Host
	protocol := "http"
	if c.Port != 80 {
		return fmt.Sprintf("%s://%s:%d", protocol, host, c.Port)
	}
	return fmt.Sprintf("%s://%s", protocol, host)
}

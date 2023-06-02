package xjwt

import "time"

type Config struct {
	Secret        string `yaml:"secret" env:"JWT_SECRET"`
	AccessExpire  int    `yaml:"access_expire" env:"JWT_ACCESS_EXPIRE"`
	RefreshExpire int    `yaml:"refresh_expire" env:"JWT_REFRESH_EXPIRE"`
}

// GetSecret
func (c *Config) GetSecret() []byte {
	return []byte(c.Secret)
}
func (c *Config) GetAccessExpire() time.Duration {
	return time.Duration(c.AccessExpire) * time.Second
}
func (c *Config) GetRefreshExpire() time.Duration {
	return time.Duration(c.RefreshExpire) * time.Second
}

func NewConfig() *Config {
	return &Config{
		Secret:        "#Secret##Secret#",
		AccessExpire:  2 * 3600,
		RefreshExpire: 360 * 24 * 3600,
	}
}

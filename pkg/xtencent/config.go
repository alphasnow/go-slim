package xtencent

import "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"

type Config struct {
	SecretId  string `yaml:"secret_id" env:"TENCENT_SECRET_ID"`
	SecretKey string `yaml:"secret_key" env:"TENCENT_SECRET_KEY"`
	Region    string `yaml:"region" env:"TENCENT_REGION"`
}

func NewConfig() *Config {
	return &Config{Region: regions.Shanghai}
}

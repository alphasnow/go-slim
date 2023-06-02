package xaliyun

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
)

func NewOpenApiConfig(cfg *Config) *openapi.Config {
	config := &openapi.Config{}
	config.SetAccessKeyId(cfg.AccessKeyId).
		SetAccessKeySecret(cfg.AccessKeySecret).
		SetRegionId(cfg.RegionId)
	return config
}

func NewRPCConfig(cfg *Config) *rpc.Config {
	config := new(rpc.Config)
	config.SetAccessKeyId(cfg.AccessKeyId).
		SetAccessKeySecret(cfg.AccessKeySecret).
		SetRegionId(cfg.RegionId)
	return config
}

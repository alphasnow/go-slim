package xtencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func NewCredential(config *Config) *common.Credential {
	return common.NewCredential(config.SecretId, config.SecretKey)
}

func NewClientProfile() *profile.ClientProfile {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 30
	//cpf.HttpProfile.Endpoint = "ft.tencentcloudapi.com"
	cpf.SignMethod = "TC3-HMAC-SHA256"
	cpf.Language = "zh-CN"

	return cpf
}

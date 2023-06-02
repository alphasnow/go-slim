package common

import (
	"go-slim/internal/models"
	"go-slim/pkg/xsignjson"
	"go-slim/pkg/xsignurl"
	"time"
)

func SignJsonRequest(data map[string]interface{}) map[string]interface{} {
	sin := xsignjson.Sign{
		Appid:    models.DefaultAppid,
		Appkey:   models.DefaultAppkey,
		NonceStr: xsignjson.RandomString(12),
	}
	result, _ := sin.Request(data)
	return result
}
func SignUrlRequest(path string) string {
	sin := xsignurl.Sign{
		Appid:  models.DefaultAppid,
		Appkey: models.DefaultAppkey,
	}
	result := sin.Signature(path, time.Now().Add(1*time.Hour).Unix())
	return result
}

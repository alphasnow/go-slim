package v2

import (
	"go-slim/api/server/controllers/v2/auth"
	"go-slim/api/server/controllers/v2/open"
	"go-slim/api/server/controllers/v2/server"
	"go-slim/api/server/controllers/v2/upload"
	"go-slim/api/server/controllers/v2/verify"
)

type Controllers struct {
	UploadRangeFile  *upload.RangeFile
	UploadSingleFile *upload.SingleFile
	UploadRule       *upload.Rules
	UploadCheck      *upload.Check

	Server *server.Server

	Paddle  *open.Paddle
	Aliyun  *open.Aliyun
	Baidu   *open.Baidu
	Tencent *open.Tencent

	AppClient    *auth.AppClient
	AuthUsername *auth.Username

	VerifyCaptcha *verify.Captcha
}

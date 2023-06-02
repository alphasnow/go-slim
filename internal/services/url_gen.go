package services

import (
	"go-slim/internal/app"
	"go-slim/internal/constants"
	"go-slim/pkg/xsignurl"
	"strings"
	"time"
)

type UrlGen struct {
	Host string
}

func NewUrlGen(config *app.Config) *UrlGen {
	return &UrlGen{Host: config.AppUrl}
}
func (u *UrlGen) Url(urlPath string) string {
	return u.Host + "/" + strings.TrimLeft(urlPath, "/")
}
func (u *UrlGen) Uploads(urlPath string) string {
	return u.Host + constants.UploadsPath + "/" + strings.TrimLeft(urlPath, "/")
}
func (u *UrlGen) Assets(urlPath string) string {
	return u.Host + constants.AssetsPath + "/" + strings.TrimLeft(urlPath, "/")
}
func (u *UrlGen) Private(urlPath string, sign *xsignurl.Sign, expire time.Time) string {
	urlPath = constants.PrivatePath + "/" + strings.TrimLeft(urlPath, "/")
	signPath := sign.Signature(urlPath, expire.Unix())
	return u.Url(signPath)
}

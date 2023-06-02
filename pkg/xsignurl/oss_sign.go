package xsignurl

import (
	"fmt"
	"net/url"
	"strconv"
)

type OssSign struct {
	AccessKeyID     string `json:"accesskeyid" form:"accesskeyid" binding:"required,max=32" validate:"required,min=8,max=32"`
	AccessKeySecret string `json:"-" url:"-"`
}

func (s *OssSign) Signature(path string, expire int64) string {
	signStr := fmt.Sprintf("%s?accesskeyid=%s&expires=%d", path, s.AccessKeyID, expire)
	signature := xMD5(signStr + "&accesskeysecret=" + s.AccessKeySecret)
	resUrl := signStr + "&signature=" + signature
	return resUrl
}

func (s *OssSign) Valid(path string, now int64) bool {
	// /tests/file.md?accesskeyid=up&expires=1658309894&signature=1dc85Iazq6XFc6ratbobIoypNTU
	reqUrl, _ := url.Parse(path)
	reqUrlQuery := reqUrl.Query()

	reqTime, _ := strconv.Atoi(reqUrlQuery.Get("expires"))
	if int(now) > reqTime {
		return false
	}

	return s.Verify(path)
}

func (s *OssSign) Verify(path string) bool {
	// /tests/file.md?accesskeyid=up&expires=1658309894&signature=1dc85Iazq6XFc6ratbobIoypNTU
	reqUrl, _ := url.Parse(path)
	reqUrlQuery := reqUrl.Query()

	signStr := fmt.Sprintf("%s?accesskeyid=%s&expires=%s", reqUrl.Path, reqUrlQuery.Get("accesskeyid"), reqUrlQuery.Get("expires"))
	signature := xMD5(signStr + "&accesskeysecret=" + s.AccessKeySecret)

	stat := signature == reqUrlQuery.Get("signature")
	return stat
}

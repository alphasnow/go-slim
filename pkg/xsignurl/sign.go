package xsignurl

import (
	"fmt"
	"net/url"
	"strconv"
)

type Sign struct {
	Appid   string `json:"appid" form:"appid" binding:"required,max=32" validate:"required,max=32"`
	Expires int    `json:"expires" form:"expires" binding:"required" validate:"required"`
	Sign    string `json:"sign" form:"sign" binding:"required,max=32" validate:"required,max=32"`
	Appkey  string `json:"-" form:"-"`
}

func (s *Sign) Signature(path string, expire int64) string {
	signStr := fmt.Sprintf("%s?appid=%s&expires=%d", path, s.Appid, expire)
	signature := xMD5(signStr + "&key=" + s.Appkey)
	resUrl := signStr + "&sign=" + signature
	return resUrl
}

func (s *Sign) valid(path string, now int64) bool {
	// /tests/file.md?appid=up&expires=1658309894&sign=1dc85Iazq6XFc6ratbobIoypNTU
	reqUrl, _ := url.Parse(path)
	reqUrlQuery := reqUrl.Query()

	reqTime, _ := strconv.Atoi(reqUrlQuery.Get("expires"))
	if int(now) > reqTime {
		return false
	}

	return s.Verify(path)
}

func (s *Sign) Parse(path string) {
	// /tests/file.md?appid=up&expires=1658309894&sign=1dc85Iazq6XFc6ratbobIoypNTU
	reqUrl, _ := url.Parse(path)
	reqUrlQuery := reqUrl.Query()
	exp, _ := strconv.Atoi(reqUrlQuery.Get("expires"))
	s.Expires = exp
	s.Sign = reqUrlQuery.Get("sign")
	s.Appid = reqUrlQuery.Get("appid")
}

func (s *Sign) Verify(path string) bool {
	// /tests/file.md?appid=up&expires=1658309894&sign=1dc85Iazq6XFc6ratbobIoypNTU
	reqUrl, _ := url.Parse(path)
	reqUrlQuery := reqUrl.Query()

	signStr := fmt.Sprintf("%s?appid=%s&expires=%s", reqUrl.Path, reqUrlQuery.Get("appid"), reqUrlQuery.Get("expires"))
	signature := xMD5(signStr + "&key=" + s.Appkey)

	stat := signature == reqUrlQuery.Get("sign")
	return stat
}

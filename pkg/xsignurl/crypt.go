package xsignurl

import (
	"crypto/md5"
	"encoding/hex"
)

func xMD5(params string) string {
	ctx := md5.New()
	ctx.Write([]byte(params))
	return hex.EncodeToString(ctx.Sum(nil))
}

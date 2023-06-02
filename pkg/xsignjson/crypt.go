package xsignjson

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func xMD5(params string) string {
	ctx := md5.New()
	ctx.Write([]byte(params))
	return hex.EncodeToString(ctx.Sum(nil))
}

func xSHA256(params string) string {
	ctx := sha256.New()
	ctx.Write([]byte(params))
	return hex.EncodeToString(ctx.Sum(nil))
}

func xHmacSHA256(params string, key string) string {
	ctx := hmac.New(sha256.New, []byte(key))
	ctx.Write([]byte(params))
	return hex.EncodeToString(ctx.Sum(nil))
}

package sign_checker

import (
	"context"
	"go-slim/pkg/xsignjson"
	"time"
)

type SignJsonChecker struct {
	SignChecker
}

func (s *SignJsonChecker) CheckNonceStr(sign *xsignjson.Sign, c context.Context) bool {
	key := "sign:noncestr:" + sign.Appid + ":" + sign.NonceStr

	val := s.Redis.Exists(c, key).Val()
	stat := val == 0
	if stat == false {
		s.Redis.Set(c, key, 1, 1*time.Minute)
	}

	return stat
}

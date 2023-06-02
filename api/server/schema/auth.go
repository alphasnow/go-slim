package schema

import "go-slim/pkg/xjwt"

type AuthLogoutReq struct {
	AccessToken string `json:"access_token"`
}
type AuthLoginRes struct {
	xjwt.IDTokensRes
	UserRes UserRes `json:"user"`
}

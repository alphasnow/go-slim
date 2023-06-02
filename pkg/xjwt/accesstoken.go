package xjwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type AccessToken struct {
	BaseToken
}

func NewAccessToken(cfg *Config) *AccessToken {
	return &AccessToken{BaseToken{Config: cfg}}
}

func (j *AccessToken) IDClaims(id int, model string) *jwt.RegisteredClaims {
	issuer := model + ":" + AccessIssuer
	return j.BaseToken.IDClaims(id, j.Config.GetAccessExpire(), issuer)
}

func (j *AccessToken) Parse(tokenString string, model string) (*jwt.RegisteredClaims, error) {
	issuer := model + ":" + AccessIssuer
	return j.BaseToken.Parse(tokenString, issuer)
}

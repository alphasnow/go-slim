package xjwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type RefreshToken struct {
	BaseToken
}

func NewRefreshToken(cfg *Config) *RefreshToken {
	return &RefreshToken{BaseToken{Config: cfg}}
}

func (j *RefreshToken) IDClaims(id int, model string) *jwt.RegisteredClaims {
	issuer := model + ":" + RefreshIssuer
	return j.BaseToken.IDClaims(id, j.Config.GetRefreshExpire(), issuer)
}

func (j *RefreshToken) Parse(tokenString string, model string) (*jwt.RegisteredClaims, error) {
	issuer := model + ":" + RefreshIssuer
	return j.BaseToken.Parse(tokenString, issuer)
}

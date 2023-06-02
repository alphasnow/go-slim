package xjwt

import "github.com/golang-jwt/jwt/v4"

type IDTokensRes struct {
	AccessToken  idTokenRes `json:"access"`
	RefreshToken idTokenRes `json:"refresh"`
}
type idTokenRes struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type IDToken struct {
	access  *AccessToken
	refresh *RefreshToken
}

func NewToken(cfg *Config) *IDToken {
	return &IDToken{access: NewAccessToken(cfg), refresh: NewRefreshToken(cfg)}
}

func (t *IDToken) Response(id int, model string) (IDTokensRes, error) {
	var resp IDTokensRes
	{
		cl := t.access.IDClaims(id, model)
		ex := cl.ExpiresAt.Unix()
		tk, err := t.access.Token(cl)
		if err != nil {
			return resp, err
		}
		resp.AccessToken = idTokenRes{Token: tk, Expire: ex}
	}
	{
		cl := t.refresh.IDClaims(id, model)
		ex := cl.ExpiresAt.Unix()
		tk, err := t.refresh.Token(cl)
		if err != nil {
			return resp, err
		}
		resp.RefreshToken = idTokenRes{Token: tk, Expire: ex}
	}

	return resp, nil
}

func (t *IDToken) GenerateAccess(id int, model string) (string, error) {
	return t.access.Token(t.access.IDClaims(id, model))
}
func (t *IDToken) ParseAccess(str string, model string) (*jwt.RegisteredClaims, error) {
	return t.access.Parse(str, model)
}

func (t *IDToken) GenerateRefresh(id int, model string) (string, error) {
	return t.access.Token(t.refresh.IDClaims(id, model))
}
func (t *IDToken) ParseRefresh(str string, model string) (*jwt.RegisteredClaims, error) {
	return t.refresh.Parse(str, model)
}

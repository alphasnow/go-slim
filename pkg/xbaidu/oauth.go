package xbaidu

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type oauth struct {
	config     *Config
	oauthToken oauthToken
	oauthError oauthError
}

type oauthToken struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int    `json:"expires_in"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	Scope         string `json:"scope"`
	SessionSecret string `json:"session_secret"`
	ExpireAt      int64  `json:"-"`
}

func (a *oauthToken) IsValid() bool {
	if a.ExpireAt == 0 || a.ExpireAt > time.Now().Unix() {
		return false
	}
	if a.AccessToken == "" {
		return false
	}
	return true
}

type oauthError struct {
	ErrorText        string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (e *oauthError) Error() string {
	return fmt.Sprintf("%s : %s", e.ErrorText, e.ErrorDescription)
}

func NewOauth(config *Config) *oauth {
	return &oauth{config: config}
}

func (a *oauth) request() (err error) {
	urlData := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {a.config.ApiKey},
		"client_secret": {a.config.SecretKey},
	}
	c := http.Client{Timeout: 10 * time.Second}
	resp, err := c.PostForm(OauthTokenUrl, urlData)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	authErr := oauthError{}
	err = json.Unmarshal(content, &authErr)
	a.oauthError = authErr
	if authErr.ErrorText != "" {
		return errors.New(authErr.Error())
	}

	token := oauthToken{}
	err = json.Unmarshal(content, &token)
	token.ExpireAt = time.Now().Unix() + int64(token.ExpiresIn)
	// 提前一分钟过期
	if token.ExpiresIn >= 600 {
		token.ExpireAt -= 60
	}
	a.oauthToken = token

	return
}
func (a *oauth) Token() (token string, err error) {
	if a.oauthToken.IsValid() == false {
		if err = a.request(); err != nil {
			return "", err
		}
	}
	return a.oauthToken.AccessToken, nil
}

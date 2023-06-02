package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-slim/api/admin/schema"
	"go-slim/internal/app"
	"go-slim/internal/constants"
	"go-slim/internal/models"
	"go-slim/pkg/xcaptcha"
	"go-slim/pkg/xhash"
	"go-slim/pkg/xjwt"
	"gorm.io/gorm"
	"net/http"
)

type AuthController struct {
	DB *gorm.DB
	// auth
	StringCaptcha *xcaptcha.StringCaptcha
	JWT           *xjwt.IDToken
	Cfg           *app.Config
}

func (s *AuthController) Login(c *gin.Context) {
	req := new(schema.LoginReq)
	if err := c.Bind(req); err != nil {
		fail(c, err)
		return
	}

	ck := s.getCaptchaKey(c)
	err := s.StringCaptcha.Verify(c, ck, req.Captcha)
	defer s.StringCaptcha.Clear(c, ck)
	if err != nil {
		fail(c, err)
		return
	}

	ad := new(models.AdminUser)
	if res := s.DB.Where("username=?", req.Username).First(ad); res.RowsAffected == 0 {
		fail(c, LoginErr)
		return
	}
	if ad.Password != xhash.MD5(req.Password) {
		fail(c, LoginErr)
		return
	}

	// jwt token
	tokenResp, err := s.JWT.Response(int(ad.ID), string(constants.JWTAdmin))
	if err != nil {
		fail(c, err)
		return
	}

	// response
	res := new(schema.LoginRes)
	res.IDTokensRes = tokenResp
	res.User = ad
	success(c, res)
}

func (s *AuthController) CaptchaImage(c *gin.Context) {
	key := s.getCaptchaKey(c)
	_, img, err := s.StringCaptcha.Generate(c, key)
	if err != nil {
		fail(c, err)
		return
	}
	c.Data(http.StatusOK, "image/png", img)
}
func (s *AuthController) CaptchaBase64(c *gin.Context) {
	key := s.getCaptchaKey(c)
	text, img, err := s.StringCaptcha.Generate(c, key)
	if err != nil {
		fail(c, err)
		return
	}
	captcha := base64.StdEncoding.EncodeToString(img)
	if s.Cfg.IsProd() == false {
		success(c, gin.H{
			"captcha": captcha,
			"text":    text,
		})
		return
	}
	success(c, gin.H{
		"captcha": captcha,
	})
}

func (s *AuthController) getCaptchaKey(c *gin.Context) string {
	ip := c.ClientIP()
	tp := c.DefaultQuery("type", "admin-login")
	key := fmt.Sprintf("%s:%s", tp, ip)
	return key
}

func (s *AuthController) Logout(c *gin.Context) {
	success(c, nil)
}

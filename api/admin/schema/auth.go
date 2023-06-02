package schema

import (
	"go-slim/internal/models"
	"go-slim/pkg/xjwt"
)

type LoginReq struct {
	Captcha  string `form:"captcha" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type LoginRes struct {
	xjwt.IDTokensRes
	User *models.AdminUser `json:"user"`
}

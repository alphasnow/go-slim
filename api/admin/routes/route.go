package routes

import (
	"github.com/gin-gonic/gin"
	"go-slim/api/admin/middlewares"
	"go-slim/pkg/xjwt"
)

type ApiAdminRoute struct {
	Gin *gin.Engine
	Ctl *Controllers

	JWT *xjwt.IDToken
}

func (s *ApiAdminRoute) Register() {

	r := s.Gin.Group("/api/admin")
	s.groupMiddleware(r)

	// auth login
	{
		_r := r.Group("/auth")
		_r.POST("login", s.Ctl.Auth.Login)
		_r.POST("logout", s.Ctl.Auth.Logout)
		_r.GET("captcha/image", s.Ctl.Auth.CaptchaImage)
		_r.GET("captcha/base64", s.Ctl.Auth.CaptchaBase64)
	}

	tk := middlewares.NewTokenChecker(s.JWT)
	r.Use(tk)

	r.POST("server/upload/image", s.Ctl.Server.UploadImage)

	AdminRouteRegister(r, s.Ctl.AdminUser, s.Ctl.AdminDepartment)
	ArticleRouteRegister(r, s.Ctl.ArticleContent, s.Ctl.ArticleCategory)
}

func (r *ApiAdminRoute) groupMiddleware(route *gin.RouterGroup) {

}

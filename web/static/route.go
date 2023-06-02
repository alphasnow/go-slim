package static

import (
	"github.com/gin-gonic/gin"
	"go-slim/internal/app"
	"go-slim/internal/config"
	"go-slim/internal/constants"
	"go-slim/web/static/middlewares"
	"go-slim/web/static/services/sign_checker"
)

type WebStaticRoute struct {
	Gin            *gin.Engine
	SignUrlChecker *sign_checker.SignUrlChecker
	Config         *config.Config
	Global         *app.Config
}

func (r *WebStaticRoute) Register() {
	// upload
	if r.Config.Http.UseNginx == false {
		//	location /upload {
		//      alias "/go-slim/release/uploads/public";
		//  }
		r.Gin.Static(constants.UploadsPath, r.Global.UploadsPath)
	}

	// assets
	if r.Config.Http.UseNginx == false {
		// https://github.com/gin-contrib
		r.Gin.StaticFile("/static/favicon.ico", r.Global.JoinRoot("web/assets/favicon.ico"))

		//	location /assets {
		//      alias "/go-slim/release/assets";
		//  }
		r.Gin.Static(constants.AssetsPath, r.Global.JoinRoot("web/assets"))
	}

	// templates
	r.Gin.LoadHTMLGlob(r.Global.JoinRoot("web/templates/**/*"))

	// download
	urlSignChecker := middlewares.NewSignUrlCheck(r.SignUrlChecker)
	download := r.Gin.Group(constants.PrivatePath, urlSignChecker)
	download.Static("", r.Global.PrivatePath)
}

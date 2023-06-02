package routes

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"go-slim/api/server/controllers/v2"
	"go-slim/api/server/middlewares"
	"go-slim/api/server/services/sign_checker"
	"go-slim/pkg/xcrypt"
	"go-slim/pkg/xlimiter"
)

type ApiV2Route struct {
	Gin *gin.Engine
	Ctl *v2.Controllers

	Limiter         *xlimiter.Limiter
	SignJsonChecker *sign_checker.SignJsonChecker
	AesCbc          *xcrypt.AesCbc
}

func (r *ApiV2Route) Register() {
	limiters := r.Limiter.NewMiddlewares([]string{"limit60:60-M", "limit10:6-M"})
	signChecker := middlewares.NewSignJsonCheck(r.SignJsonChecker)

	api1 := r.Gin.Group("/api/v2")
	r.groupMiddleware(api1)

	upload := api1.Group("/upload")
	{
		tokenChecker := middlewares.NewUploadTokenCheck(r.AesCbc)
		upload.POST("/rule", signChecker, r.Ctl.UploadRule.GetRule)
		upload.POST("/check", signChecker, r.Ctl.UploadCheck.Prepare)
		upload.POST("/range", tokenChecker, r.Ctl.UploadRangeFile.Upload)
		upload.POST("/single", tokenChecker, r.Ctl.UploadSingleFile.Upload)
	}

	server := api1.Group("/server")
	{
		server.GET("timestamp", limiters["limit60:60-M"], r.Ctl.Server.Timestamp)
		server.GET("health", limiters["limit10:6-M"], r.Ctl.Server.Health)
	}

	auth := api1.Group("/auth", signChecker)
	{
		auth.POST("/client/register", r.Ctl.AppClient.Register)
		auth.POST("/username/register", r.Ctl.AuthUsername.Register)
	}

	verify := api1.Group("/verify", signChecker)
	{
		verify.POST("/captcha/image", r.Ctl.VerifyCaptcha.Image)
	}

	open := api1.Group("/open", signChecker)
	{
		open.POST("/human-segment", r.Ctl.Paddle.HumanSegment)
		open.POST("/body-segment", r.Ctl.Baidu.BodySeg)
		open.POST("/change-age", r.Ctl.Tencent.ChangeAge)
		open.POST("/common-segment", r.Ctl.Aliyun.CommonSeg)
	}

}

func (r *ApiV2Route) groupMiddleware(route *gin.RouterGroup) {
	// REQUEST-ID
	route.Use(requestid.New())
}

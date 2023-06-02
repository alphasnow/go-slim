package routes

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	limits "github.com/gin-contrib/size"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	adminRoutes "go-slim/api/admin/routes"
	"go-slim/api/server/routes"
	"go-slim/api/server/schema"
	"go-slim/internal/config"
	"go-slim/internal/middlewares"
	"go-slim/pkg/xlog"
	"go-slim/web/home"
	"go-slim/web/static"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

type Route interface {
	Register()
}

type Routes struct {
	Gin        *gin.Engine
	Config     *config.Config
	Log        *zap.Logger
	LogManager *xlog.Manager

	ApiAdminRoute *adminRoutes.ApiAdminRoute
	ApiV1Route    *routes.ApiV1Route
	ApiV2Route    *routes.ApiV2Route

	WebStaticRoute *static.WebStaticRoute
	WebHomeRoute   *home.WebHomeRoute
}

var _ Route = (*Routes)(nil)

func (r *Routes) Register() {
	r.WebStaticRoute.Register()

	if r.Config.Http.UseNginx == false {
		r.globalMiddlewareNotUseNginx()
	}

	r.globalErrorMiddleware()

	r.WebHomeRoute.Register()

	if r.Config.App.IsDev() {
		// swagger
		r.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		// pprof
		//	// default is "debug/pprof"
		//	pprof.Register(r.Gin)
	}

	// cors
	r.Gin.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	r.ApiV1Route.Register()
	r.ApiV2Route.Register()

	r.ApiAdminRoute.Register()
}

func (r *Routes) globalErrorMiddleware() {

	r.Gin.Use(gin.CustomRecoveryWithWriter(&errorLogger{r.Log}, middlewares.CustomRecoveryHandler(r.Log)))
	r.Gin.Use(middlewares.ErrorHandler(r.Log))

	// 404 405
	r.Gin.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		// module: admin
		if strings.HasPrefix(path, "/api/admin") {
			adminRoutes.HandleNoRoute(c)
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.NotFound))
	})
	r.Gin.NoMethod(func(c *gin.Context) {
		path := c.Request.URL.Path
		// module: admin
		if strings.HasPrefix(path, "/api/admin") {
			adminRoutes.HandleNoMethod(c)
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.MethodNotAllowed))
	})
}

func (r *Routes) globalMiddlewareNotUseNginx() {
	// GZIP
	r.Gin.Use(gzip.Gzip(gzip.BestSpeed))

	// POST LIMIT
	if r.Config.Http.MaxPostSize != 0 {
		r.Gin.Use(limits.RequestSizeLimiter(int64(r.Config.Http.MaxPostSize)))
	}

	// AccessLog
	r.Gin.Use(ginzap.Ginzap(r.Log, "2006-01-02 15:04:05", true))
}

type errorLogger struct {
	*zap.Logger
}

func (l *errorLogger) Write(p []byte) (n int, err error) {
	str := string(p)
	fmt.Println(str)
	// "[0m"
	// str = str[7 : len(str)-5]
	str = str[60 : len(str)-5]
	l.Error(str)
	return len(str), nil
}

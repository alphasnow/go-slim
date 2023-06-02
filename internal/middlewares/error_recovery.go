package middlewares

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
)

func CustomRecoveryHandler(logger *zap.Logger) gin.RecoveryFunc {
	return func(c *gin.Context, err any) {

		logger.Error("[Panic]",
			zap.Any("error", err),
		)
		logger.Error(string(debug.Stack()))

		reqID := requestid.Get(c)

		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"err_code": 500,
			"err_msg":  err,
			"req_id":   reqID,
		})
	}
}

// ErrorHandler
// refer: https://zhuanlan.zhihu.com/p/357460091
// refer: https://github.com/gin-gonic/gin/issues/342#issuecomment-111772927
// example: c.AbortWithError(200, errors.New("diy error msg"))
// example: c.Error(errors.New("some error message"))
func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {

			if length := len(c.Errors); length > 0 {

				logger.Error("[Error]",
					zap.Any("errors", c.Errors),
				)
				logger.Error(string(debug.Stack()))

				errMsg := ""
				for _, v := range c.Errors {
					errMsg += v.Error() + ";"
				}
				reqID := requestid.Get(c)
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"err_code": 500,
					"err_msg":  errMsg,
					"req_id":   reqID,
				})

			}
		}()
		c.Next()
	}
}

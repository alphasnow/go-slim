package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-slim/api/admin/schema"
	"go-slim/internal/constants"
	"go-slim/pkg/xjwt"
	"net/http"
	"strconv"
)

func NewTokenChecker(jwt *xjwt.IDToken) gin.HandlerFunc {
	return func(c *gin.Context) {
		tk := getTokenStr(c)
		fail := func(err string) {
			c.AbortWithStatusJSON(http.StatusOK, schema.Response{
				Success:      false,
				ErrorMessage: err,
				ErrorCode:    401,
				ShowType:     schema.REDIRECT,
			})
		}
		if tk == "" {
			fail("401")
			return
		}
		claims, err := jwt.ParseAccess(tk, string(constants.JWTAdmin))
		if err != nil {
			fail("401" + err.Error())
			return
		}

		c.Set("user_id", claims.Subject)
		c.Next()
	}
}

func getTokenStr(c *gin.Context) string {
	var reqToken string
	reqToken = c.GetHeader("Authorization")
	// Authorization: Bearer
	if reqToken != "" && reqToken[7:] != "" {
		return reqToken[7:]
	}
	return ""
}

func GetTokenUserID(c *gin.Context) int {
	userID, exist := c.Get("user_id")
	if exist == false {
		return 0
	}
	id, err := strconv.Atoi(userID.(string))
	if err != nil {
		return 0
	}
	return id
}

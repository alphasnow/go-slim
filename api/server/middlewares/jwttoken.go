package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-slim/api/server/schema"
	"net/http"
)

const (
	TokenField  = "token"
	TokenHeader = "Authorization"

	UserID  = "user_id" // c.Set("token_id", id)
	TokenID = "token_id"
)

func GetTokenByRequest(c *gin.Context) string {
	var reqToken string
	reqToken = c.GetHeader(TokenHeader)
	// Authorization: Bearer
	if reqToken != "" && reqToken[7:] != "" {
		return reqToken[7:]
	}
	reqToken, _ = c.GetQuery(TokenField)
	if reqToken != "" {
		return reqToken
	}
	reqToken, _ = c.GetPostForm(TokenField)
	if reqToken != "" {
		return reqToken
	}
	reqToken, _ = c.Cookie(TokenField)
	if reqToken != "" {
		return reqToken
	}
	return ""
}

type MiddlewareToken interface {
	Parse(string) (*jwt.RegisteredClaims, error)
	ParseID(*jwt.RegisteredClaims) (int, error)
}

func NewTokenMiddleware(mt MiddlewareToken) gin.HandlerFunc {
	return func(c *gin.Context) {
		tk := GetTokenByRequest(c)
		if tk == "" {
			c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.Unauthorized))
			return
		}

		claims, err := mt.Parse(tk)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.Unauthorized, err))
			return
		}
		id, _ := mt.ParseID(claims)

		c.Set(UserID, id)
		c.Set(TokenID, claims.ID)

		c.Next()
	}
}

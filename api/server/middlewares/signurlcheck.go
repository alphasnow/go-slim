package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-slim/api/server/schema"
	"go-slim/api/server/services/sign_checker"
	"go-slim/pkg/xsignurl"
	"net/http"
	"strconv"
	"time"
)

const SignUrlExpireKey = "expires"
const SignUrl = "_internal/middlewares/sign_url"

func NewSignUrlCheck(s *sign_checker.SignUrlChecker) gin.HandlerFunc {

	return func(c *gin.Context) {
		sign := &xsignurl.Sign{}
		if err := c.ShouldBind(sign); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.SignUrlParam, err))
			return
		}

		expires, _ := strconv.Atoi(c.Query(SignUrlExpireKey))
		if expires < int(time.Now().Unix()) {
			c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.SignUrlExpire))
			return
		}

		app, err := s.FindApp(sign.Appid, c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.SignUrlAppid))
			return
		}
		sign.Appkey = app.Appkey

		status := sign.Verify(c.Request.URL.String())
		if status == false {
			c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.SignUrlSign))
			return
		}

		c.Set(AppModel, app)
		c.Set(SignUrl, sign)
		c.Next()
	}
}

func GetSignUrl(c *gin.Context) (*xsignurl.Sign, error) {
	sign, exist := c.Get(SignUrl)
	if exist == false {
		return nil, errors.New("sign_url not exist")
	}
	signObj, ok := sign.(*xsignurl.Sign)
	if ok == false {
		return nil, errors.New("sign_url type error")
	}
	return signObj, nil
}

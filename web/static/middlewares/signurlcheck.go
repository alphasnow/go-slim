package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-slim/pkg/xsignurl"
	"go-slim/web/static/services/sign_checker"
	"net/http"
	"strconv"
	"time"
)

const AppModel = "_internal/middlewares/app_model"
const SignUrlExpireKey = "expires"
const SignUrl = "_internal/middlewares/sign_url"

func NewSignUrlCheck(s *sign_checker.SignUrlChecker) gin.HandlerFunc {
	return func(c *gin.Context) {
		sign := &xsignurl.Sign{}
		if err := c.ShouldBind(sign); err != nil {
			c.Abort()
			c.String(http.StatusOK, "params error")
			return
		}

		expires, _ := strconv.Atoi(c.Query(SignUrlExpireKey))
		if expires < int(time.Now().Unix()) {
			c.Abort()
			c.String(http.StatusOK, "time expire")
			return
		}

		app, err := s.FindApp(sign.Appid, c)
		if err != nil {
			c.Abort()
			c.String(http.StatusOK, "APPID error")
			return
		}
		sign.Appkey = app.Appkey

		status := sign.Verify(c.Request.URL.String())
		if status == false {
			c.Abort()
			c.String(http.StatusOK, "sign error")
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

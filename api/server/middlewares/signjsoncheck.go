package middlewares

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/server/schema"
	"go-slim/api/server/services/sign_checker"
	"go-slim/pkg/xsignjson"
	"net/http"
)

const AppModel = "_internal/middlewares/app_model"
const BodySign = "_internal/middlewares/sign_json"

func NewSignJsonCheck(s *sign_checker.SignJsonChecker) gin.HandlerFunc {
	return func(c *gin.Context) {
		sign := &xsignjson.Sign{}
		if err := c.ShouldBindBodyWith(sign, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.SignJsonParam, err))
			return
		}

		if s.CheckNonceStr(sign, c) == false {
			c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.SignJsonNonceStr))
			return
		}

		app, err := s.FindApp(sign.Appid, c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.SignJsonAppid, err))
			return
		}
		sign.Appkey = app.Appkey

		bd, _ := c.Get(gin.BodyBytesKey)
		bdMap := map[string]interface{}{}
		if err := json.Unmarshal(bd.([]byte), &bdMap); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.SignJsonBody, err))
			return
		}

		status := sign.Verify(bdMap)
		if status == false {
			c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.SignJsonSign))
			return
		}

		c.Set(AppModel, app)
		c.Set(BodySign, sign)
		c.Next()
	}
}

func GetSignJson(c *gin.Context) (*xsignjson.Sign, error) {
	sign, exist := c.Get(BodySign)
	if exist == false {
		return nil, errors.New("sign_json not exist")
	}
	signObj, ok := sign.(*xsignjson.Sign)
	if ok == false {
		return nil, errors.New("sign_json type error")
	}
	return signObj, nil
}

func GetAppCache(c *gin.Context) (*sign_checker.AppCache, error) {
	sign, exist := c.Get(AppModel)
	if exist == false {
		return nil, errors.New("app model not exist")
	}
	signObj, ok := sign.(*sign_checker.AppCache)
	if ok == false {
		return nil, errors.New("app model type error")
	}
	return signObj, nil
}

package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/middlewares"
	"go-slim/api/server/schema"
	"go-slim/internal/models"
	"go-slim/pkg/xcaptcha"
	"go-slim/pkg/xjwt"
	"gorm.io/gorm"
)

type Username struct {
	DB            *gorm.DB
	JWT           *xjwt.IDToken
	StringCaptcha *xcaptcha.StringCaptcha
}

func (s *Username) Register(c *gin.Context) {
	req := schema.AuthUsernameLoginReq{}
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	// verify
	key := req.DeviceUUID + ":" + c.Request.URL.Path
	err := s.StringCaptcha.Verify(c, key, req.VerifyCode)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}
	defer s.StringCaptcha.Clear(c, key)

	user := &models.User{}
	s.DB.First(user, "username=?", req.Username)
	if user.ID != 0 {
		common.Error(c, schema.BadRequest, errors.New("username have been used"))
		return
	}

	app, _ := middlewares.GetAppCache(c)

	user.Username = req.Username
	user.AppID = app.ID
	user.Password = req.Password
	s.DB.Create(user)
}

func (s *Username) Login(c *gin.Context) {
	// todo: ...
}

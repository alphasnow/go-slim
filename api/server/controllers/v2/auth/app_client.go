package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/logic"
	"go-slim/api/server/middlewares"
	"go-slim/api/server/schema"
	"go-slim/internal/constants"
	"go-slim/internal/models"
	"go-slim/pkg/xjwt"
	"gorm.io/gorm"
)

type AppClient struct {
	DB         *gorm.DB
	JWT        *xjwt.IDToken
	UserClient *logic.UserClient
}

func (s *AppClient) Register(c *gin.Context) {
	req := &schema.AppClientRegisterReq{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	app, err := middlewares.GetAppCache(c)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	var uc *models.UserClient
	uc, err = s.UserClient.FirstByRequest(c, req, app.ID)
	if err != nil {
		uc, err = s.UserClient.CreateByRequest(c, req, app.ID)
		if err != nil {
			common.Error(c, schema.BadRequest, err)
			return
		}
	}

	tokenResp, err := s.JWT.Response(int(uc.User.ID), string(constants.JWTUser))
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	var resp schema.AppClientRegisterRes
	resp.IDTokensRes = tokenResp
	resp.UserRes = schema.UserRes{UID: uc.User.UID, CreatedAt: uc.User.CreatedAt.Unix()}
	resp.ClientRes = schema.ClientRes{DeviceID: uc.DeviceID, DeviceUUID: uc.DeviceUUID}

	common.Success(c, resp)
}

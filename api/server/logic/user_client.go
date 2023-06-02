package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/pjebs/optimus-go"
	"go-slim/api/server/schema"
	"go-slim/internal/models"
	"go-slim/pkg/xuuid"
	"gorm.io/gorm"
	"time"
)

type UserClient struct {
	DB      *gorm.DB
	Optimus *optimus.Optimus
}

func (s *UserClient) CreateByRequest(c *gin.Context, req *schema.AppClientRegisterReq, appID uint) (*models.UserClient, error) {
	u := &models.User{}
	u.CreatedAt = time.Now()
	u.AppID = appID

	uc := &models.UserClient{}
	uc.UserAgent = c.GetHeader("User-Agent")
	uc.AcceptLanguage = c.GetHeader("Accept-Language")
	uc.ClientIP = c.ClientIP()
	uc.DeviceUUID = xuuid.Generate()
	uc.DeviceID = req.DeviceID
	uc.AppVersion = req.AppVersion
	uc.AppID = appID
	uc.User = u

	// https://gorm.io/zh_CN/docs/associations.html#Association-Mode
	// https://gorm.io/zh_CN/docs/transactions.html
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		if res := s.DB.Create(u); res.Error != nil {
			return res.Error
		}

		uid := uint(s.Optimus.Encode(uint64(u.ID)))
		if res := s.DB.Model(u).Update("uid", uid); res.Error != nil {
			return res.Error
		}

		if res := s.DB.Save(uc); res.Error != nil {
			return res.Error
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return uc, nil
}

func (s *UserClient) FirstByRequest(c *gin.Context, req *schema.AppClientRegisterReq, appID uint) (*models.UserClient, error) {
	uc := &models.UserClient{}

	dbRes := s.DB.WithContext(c).Preload("User").First(uc, "device_id=? and app_id=?", req.DeviceID, appID)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return uc, nil
}

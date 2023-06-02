package sign_checker

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"go-slim/internal/models"
	"gorm.io/gorm"
	"time"
)

type SignChecker struct {
	Redis *redis.Client
	Mysql *gorm.DB
}

func (s *SignChecker) FindAppkey(appid string, c context.Context) (string, error) {
	key := "sign:appkey:" + appid
	rdRes := s.Redis.Get(c, key).Val()
	if rdRes == "0" {
		return "", errors.New("appkey not found")
	}
	if rdRes != "" {
		return rdRes, nil
	}

	app := &models.App{}
	if dbRes := s.Mysql.Model(models.App{}).Where("appid=?", appid).First(app); dbRes.Error != nil {
		s.Redis.Set(c, key, "0", 10*time.Minute)
		return "", dbRes.Error
	}

	s.Redis.Set(c, key, app.Appkey, 60*time.Minute)
	return app.Appkey, nil
}

func (s *SignChecker) FindApp(appid string, c context.Context) (*AppCache, error) {
	// redis key
	key := "sign:appkey:" + appid
	appConvert := &AppConvert{}

	rdRes := s.Redis.HGetAll(c, key).Val()
	aid, ok := rdRes["appkey"]
	if ok == true {
		if aid != "0" {
			app := appConvert.ToAppModel(rdRes)
			return &app, nil
		}
		return nil, errors.New("appkey not found")
	}

	app := models.App{}
	if dbRes := s.Mysql.Model(models.App{}).Where("appid=?", appid).First(&app); dbRes.Error != nil {
		appCache := AppCache{Appid: appid, ID: 0, Appkey: "0"}
		appMap := appConvert.ToMapString(appCache)
		s.Redis.HSet(c, key, appMap)
		s.Redis.Expire(c, key, 10*time.Minute)
		return nil, dbRes.Error
	}

	appCache := AppCache{Appid: app.Appid, ID: app.ID, Appkey: app.Appkey}
	appMap := appConvert.ToMapString(appCache)
	s.Redis.HSet(c, key, appMap)
	s.Redis.Expire(c, key, 60*time.Minute)

	return &appCache, nil
}

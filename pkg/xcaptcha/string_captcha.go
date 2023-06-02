package xcaptcha

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"strings"
	"time"
)

const RedisPrefix = "captcha"
const RedisExpire = 10 * time.Minute

type StringCaptcha struct {
	driver *base64Captcha.DriverString
	redis  *redis.Client
}

func NewStringCaptcha(redis *redis.Client) *StringCaptcha {
	driver := base64Captcha.NewDriverString(60, 200, 0, base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSineLine, 4, "123456789abcdefghjklmnpqrstuvwxyz", &color.RGBA{255, 255, 255, 0}, nil, []string{"chromohv.ttf"})
	return &StringCaptcha{driver: driver, redis: redis}
}

func (s *StringCaptcha) Generate(ctx context.Context, key string) (text string, img []byte, err error) {
	text = base64Captcha.RandText(s.driver.Length, s.driver.Source)

	item, err := s.driver.DrawCaptcha(text)
	if err != nil {
		return text, nil, err
	}
	itemChar := item.(*base64Captcha.ItemChar)
	img = itemChar.BinaryEncoding()

	if res := s.redis.Set(ctx, RedisPrefix+":"+key, text, RedisExpire); res.Err() != nil {
		return text, img, res.Err()
	}

	return
}

func (s *StringCaptcha) Verify(ctx context.Context, key string, req string) (err error) {
	res := s.redis.Get(ctx, RedisPrefix+":"+key)
	if res.Val() == "" {
		return errors.New("captcha expire")
	}
	if res.Err() != nil {
		return res.Err()
	}

	if res.Val() != strings.ToLower(strings.TrimSpace(req)) {
		return errors.New("captcha error")
	}
	return
}

func (s *StringCaptcha) Clear(ctx context.Context, key string) error {
	if res := s.redis.Del(ctx, RedisPrefix+":"+key); res.Err() != nil {
		return res.Err()
	}

	return nil
}

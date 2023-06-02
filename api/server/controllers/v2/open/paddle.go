package open

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/schema"
	"go-slim/pkg/ximage"
	"go-slim/pkg/xpaddle"
	"image"
	"io/ioutil"
	"mime/multipart"
)

type Paddle struct {
	Humanseg *xpaddle.HumansegMobileClient
}

func (s *Paddle) HumanSegment(c *gin.Context) {
	fh, err := c.FormFile("image")
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	err = s.checkFile(fh)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	srcBytes, err := s.getFileBytes(fh)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	srcBytes, err = ximage.Convert(srcBytes, ximage.JPEG)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	maskBase64, err := s.humansegMobile(srcBytes)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	pngBase64, err := s.compositeMask(srcBytes, maskBase64)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	common.Success(c, schema.ImageRes{Image: pngBase64})
	return
}

func (s Paddle) checkFile(file *multipart.FileHeader) error {
	if file.Size > 5*1024*1024 {
		return errors.New("file size gt 5M")
	}

	f, err := file.Open()
	cfg, ext, err := image.DecodeConfig(f)
	if err != nil {
		return err
	}

	exts := map[string]bool{"jpeg": true, "png": true, "gif": true}
	if _, ok := exts[ext]; ok == false {
		return errors.New("file type invalid")
	}

	if cfg.Width > 2000 || cfg.Height > 2000 {
		return errors.New("file size gt 2000")
	}

	rate := float32(cfg.Height / cfg.Width)
	if rate > 4 || rate < 0.25 {
		return errors.New("file rate gt 4")
	}

	return nil
}

func (s Paddle) getFileBytes(file *multipart.FileHeader) ([]byte, error) {
	fp, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	srcBytes, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}
	return srcBytes, nil
}

func (s *Paddle) humansegMobile(srcBytes []byte) (string, error) {
	imageBase64 := base64.StdEncoding.EncodeToString(srcBytes)

	resp, err := s.Humanseg.Request([]string{imageBase64})
	if err != nil {
		return "", err
	}
	return resp[0], nil
}

func (s *Paddle) compositeMask(srcBytes []byte, maskBase64 string) (string, error) {
	maskBytes, err := base64.StdEncoding.DecodeString(maskBase64)
	if err != nil {
		return "", err
	}

	rsrByte, err := ximage.ByteMaskComposite(srcBytes, maskBytes)
	if err != nil {
		return "", err
	}

	rstBase64 := base64.StdEncoding.EncodeToString(rsrByte)
	return rstBase64, nil
}

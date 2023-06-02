package image_tool

import (
	"bytes"
	"encoding/base64"
	"go-slim/api/server/schema"
	"image"
)

type ImageInformation struct {
}

func (s *ImageInformation) ByBinary(imInfo *schema.ImageInfo, imgBinary []byte) error {
	imgReader := bytes.NewReader(imgBinary)
	imInfo.Size = imgReader.Size()

	imgCnf, ext, err := image.DecodeConfig(imgReader)
	if err != nil {
		return err
	}
	imInfo.Width = imgCnf.Width
	imInfo.Height = imgCnf.Height
	imInfo.Ext = ext

	return nil
}

func (s *ImageInformation) ByBase64(imInfo *schema.ImageInfo, imgBase64 string) error {
	imInfo.Base64 = int64(len(imgBase64))

	imgBinary, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return err
	}

	return s.ByBinary(imInfo, imgBinary)
}

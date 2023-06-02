package image_tool

import (
	"errors"
	"github.com/samber/lo"
	"go-slim/api/server/schema"
)

type ImageChecker struct {
}

func (s *ImageChecker) CheckRules(imgInfo *schema.ImageInfo, imgRules *schema.ImageRule) error {
	if imgRules.Base64 > 0 && imgInfo.Base64 > 0 && imgInfo.Base64 > imgRules.Base64 {
		return errors.New("image base64 size error")
	}

	if imgRules.Size > 0 && imgInfo.Size > imgRules.Size {
		return errors.New("image size error")
	}

	if len(imgRules.Ext) > 0 {
		// https://github.com/samber/lo#contains
		if lo.Contains[string](imgRules.Ext, imgInfo.Ext) == false {
			return errors.New("image ext error")
		}
	}

	if imgRules.Width > 0 && imgInfo.Width > imgRules.Width {
		return errors.New("image max width error")
	}
	if imgRules.Height > 0 && imgInfo.Height > imgRules.Height {
		return errors.New("image max height error")
	}
	if imgRules.MinWidth > 0 && imgInfo.Width < imgRules.MinWidth {
		return errors.New("image min width error")
	}
	if imgRules.MinHeight > 0 && imgInfo.Height < imgRules.MinHeight {
		return errors.New("image min height error")
	}

	if imgRules.Ratio > 0 {
		var long, short int
		if imgInfo.Height > imgInfo.Width {
			long, short = imgInfo.Height, imgInfo.Width
		} else {
			short, long = imgInfo.Height, imgInfo.Width
		}
		ratio := float32(long / short)
		if ratio > imgRules.Ratio {
			return errors.New("image ratio error")
		}
	}

	return nil
}

package open

import (
	"bytes"
	"encoding/base64"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/schema"
	"go-slim/api/server/services/image_tool"
	"go-slim/pkg/xaliyun"
)

type Aliyun struct {
	Seg *xaliyun.ImageSeg
	Oss *oss.Bucket
}
type commonSegReq struct {
	Image string `json:"image" binding:"required"`
	Form  string `json:"form"`
}
type commonSegRes struct {
	Image *string `json:"image"`
}

// https://help.aliyun.com/document_detail/151960.html
var commonSegImageRule = schema.ImageRule{
	Ext:       []string{"jpeg", "jpg", "png", "bmp", "webp"},
	Size:      3 * 1024 * 1024,
	MinWidth:  32,
	MinHeight: 32,
	Width:     2000,
	Height:    2000,
}

// CommonSeg
func (s *Aliyun) CommonSeg(c *gin.Context) {
	var err error

	req := new(commonSegReq)
	if err = c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	imInf := new(schema.ImageInfo)
	imInform := new(image_tool.ImageInformation)
	if err = imInform.ByBase64(imInf, req.Image); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}
	imChecker := new(image_tool.ImageChecker)
	if err = imChecker.CheckRules(imInf, &commonSegImageRule); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	var imgBinary []byte
	imgBinary, err = base64.StdEncoding.DecodeString(req.Image)
	ossName := requestid.Get(c) + "." + imInf.Ext
	ossPath := "uploads/" + ossName
	opts := []oss.Option{
		oss.ObjectACL(oss.ACLPrivate),
	}
	err = s.Oss.PutObject(ossPath, bytes.NewReader(imgBinary), opts...)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}
	ossUrl, err := s.Oss.SignURL(ossPath, oss.HTTPGet, 600)
	if err != nil {
		// _ = s.Oss.DeleteObject(ossPath)
		common.Error(c, schema.BadRequest, err)
		return
	}

	resp, err := s.Seg.SegmentCommonImage(ossUrl, req.Form)
	_ = s.Oss.DeleteObject(ossPath)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}
	convert := new(image_tool.ImageConverter)
	urlBase64, err := convert.UrlToBase64(*resp.Body.Data.ImageURL)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	common.Success(c, commonSegRes{
		Image: urlBase64,
	})
}

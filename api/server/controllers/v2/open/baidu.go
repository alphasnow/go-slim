package open

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/schema"
	"go-slim/api/server/services/image_tool"
	"go-slim/pkg/xbaidu"
)

type Baidu struct {
	Client *xbaidu.Client
}

type bodySegReq struct {
	Image string `json:"image" binding:"required"`
	Type  string `json:"type"`
}
type bodySegRes struct {
	Labelmap   *string `json:"labelmap"`
	Scoremap   *string `json:"scoremap"`
	Foreground *string `json:"foreground"`
}

var bodySegRule = schema.ImageRule{
	Ext:       []string{"jpeg", "jpg", "png", "bmp"},
	Base64:    4 * 1024 * 1024,
	MinWidth:  50,
	MinHeight: 50,
	Width:     4096,
	Height:    4096,
}

func (s *Baidu) BodySeg(c *gin.Context) {
	var err error

	req := new(bodySegReq)
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
	if err = imChecker.CheckRules(imInf, &bodySegRule); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	resp, err := s.Client.BodySeg(req.Image, req.Type)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	common.Success(c, bodySegRes{
		Labelmap:   &resp.Labelmap,
		Foreground: &resp.Foreground,
		Scoremap:   &resp.Scoremap,
	})
}

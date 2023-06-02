package open

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/schema"
	"go-slim/api/server/services/image_tool"
	"go-slim/pkg/xtencent"
)

type Tencent struct {
	Client *xtencent.FtClient
}
type changeAgeReq struct {
	Image string `json:"image" binding:"required"`
	Ages  []int  `json:"ages"`
}
type changeAgeRes struct {
	Image *string `json:"image"`
}

var changeAgeImageRule = schema.ImageRule{
	Ext:    []string{"jpeg", "jpg", "png", "bmp"},
	Base64: 5 * 1024 * 1024,
}

// ChangeAge
// https://cloud.tencent.com/document/product/1202/41968
func (s *Tencent) ChangeAge(c *gin.Context) {
	var err error

	req := new(changeAgeReq)
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
	if err = imChecker.CheckRules(imInf, &changeAgeImageRule); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	resp, err := s.Client.ChangeAgePic(&req.Image, req.Ages)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	common.Success(c, changeAgeRes{Image: resp.Response.ResultImage})
}

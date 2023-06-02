package xaliyun

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	imageseg "github.com/alibabacloud-go/imageseg-20191230/v2/client"
)

// ImageSeg
// @Link https://help.aliyun.com/document_detail/151960.html
type ImageSeg struct {
	config *openapi.Config
}

func NewImageSeg(cfg *openapi.Config) *ImageSeg {
	return &ImageSeg{config: cfg}
}

func (s *ImageSeg) SegmentCommonImage(imageUrl string, returnForm string) (*imageseg.SegmentCommonImageResponse, error) {
	c, err := imageseg.NewClient(s.config)
	if err != nil {
		return nil, err
	}
	req := new(imageseg.SegmentCommonImageRequest)
	req.SetImageURL(imageUrl)
	req.SetReturnForm(returnForm)
	resp, err := c.SegmentCommonImage(req)
	return resp, err
}

package xtencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ft "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ft/v20200304"
)

type FtClient struct {
	client *ft.Client
}

func NewFtClient(config *Config, credential *common.Credential, clientProfile *profile.ClientProfile) *FtClient {
	c, _ := ft.NewClient(credential, config.Region, clientProfile)

	return &FtClient{client: c}
}

func (c *FtClient) Client() *ft.Client {
	return c.client
}

func (c *FtClient) ChangeAgePic(image *string, ages []int) (*ft.ChangeAgePicResponse, error) {
	req := ft.NewChangeAgePicRequest()
	req.Image = image
	infos := make([]*ft.AgeInfo, len(ages))
	for k, v := range ages {
		age := int64(v)
		infos[k] = &ft.AgeInfo{Age: &age}
	}
	req.AgeInfos = infos
	rspType := "base64"
	req.RspImgType = &rspType

	res, err := c.client.ChangeAgePic(req)
	return res, err
}

package xpaddle

import (
	"errors"
)

type HumansegMobileClient struct {
	Client *Client
}

func (c HumansegMobileClient) Request(imageBase64 []string) ([]string, error) {
	req := humansegMobileReq{Images: imageBase64}
	resp := &humansegMobileResp{}

	err := c.Client.Request("humanseg_mobile", req, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() == false {
		return nil, errors.New(resp.Error())
	}

	return resp.Images(), nil
}

type humansegMobileReq struct {
	Images []string `json:"images"`
}
type humansegMobileResp struct {
	ServerResp
	Results []struct {
		Data string `json:"data"`
	} `json:"results"`
}

func (resp *humansegMobileResp) Images() []string {
	images := make([]string, len(resp.Results))
	for k, v := range resp.Results {
		images[k] = v.Data
	}
	return images
}

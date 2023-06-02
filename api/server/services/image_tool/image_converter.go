package image_tool

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

type ImageConverter struct {
}

func (s *ImageConverter) UrlToBase64(imageUrl string) (*string, error) {
	urlResp, err := http.Get(imageUrl)
	if err != nil {
		return nil, err
	}
	urlCon, err := ioutil.ReadAll(urlResp.Body)
	if err != nil {
		return nil, err
	}
	urlBase64 := base64.StdEncoding.EncodeToString(urlCon)
	return &urlBase64, nil
}

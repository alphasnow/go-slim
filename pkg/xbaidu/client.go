package xbaidu

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	oauth *oauth
}

func NewClient(oauth *oauth) *Client {
	return &Client{oauth: oauth}
}

func (c *Client) PostForm(aipUrl string, urlData url.Values, aipResp interface{}) (err error) {
	token, err := c.oauth.Token()
	if err != nil {
		return err
	}
	// https://ai.baidu.com/ai-doc/BODY/Fk3cpyxua
	hc := http.Client{Timeout: 10 * time.Second}
	tokenQuery := fmt.Sprintf("?access_token=%s", token)
	resp, err := hc.PostForm(aipUrl+tokenQuery, urlData)
	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &aipResp)
	return
}
func (c *Client) PostFormWithContext(ctx context.Context, aipUrl string, urlData url.Values, aipResp interface{}) (err error) {

	token, err := c.oauth.Token()
	if err != nil {
		return err
	}
	tokenQuery := fmt.Sprintf("?access_token=%s", token)

	body := strings.NewReader(urlData.Encode())
	req, err := http.NewRequest(http.MethodPost, aipUrl+tokenQuery, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.WithContext(ctx)

	hc := http.Client{Timeout: 10 * time.Second}
	resp, err := hc.Do(req)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &aipResp)
	return
}

func (c *Client) BodySeg(image string, types string) (result *bodySegResp, err error) {
	// https://ai.baidu.com/ai-doc/BODY/Fk3cpyxua
	postData := url.Values{"image": {image}, "type": {types}}
	resp := &bodySegResp{}
	if err = c.PostForm(BodySegUrl, postData, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

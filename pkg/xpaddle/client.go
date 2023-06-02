package xpaddle

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Client struct {
	Url string // http://127.0.0.1:8866/predict
}

func NewClient(cfg *Config) *Client {
	return &Client{Url: cfg.Url}
}

func (c *Client) Request(module string, data interface{}, res interface{}) error {
	url := c.Url + "/" + module
	jsonBytes, _ := json.Marshal(data)
	jsonReader := bytes.NewReader(jsonBytes)
	req, _ := http.NewRequest(http.MethodPost, url, jsonReader)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Content-Length", strconv.Itoa(jsonReader.Len()))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(respBody, res)
}

package common

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type HttpClient struct {
	Router *gin.Engine
	query  url.Values
	header map[string]string
}

func (c *HttpClient) Request(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	// query
	if c.query != nil {
		urlQuery := c.query.Encode()
		if urlQuery != "" {
			path += "?" + urlQuery
		}
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, path, body)

	// header
	for k, v := range c.header {
		r.Header.Set(k, v)
	}

	c.Router.ServeHTTP(w, r)

	// reset
	c.query = url.Values{}
	c.header = map[string]string{}

	return w
}

func (c *HttpClient) WithHeader(header map[string]string) *HttpClient {
	if c.header == nil {
		c.header = header
	}
	for k, v := range header {
		c.header[k] = v
	}
	return c
}

func (c *HttpClient) WithQuery(query map[string]string) *HttpClient {
	if c.query == nil {
		c.query = url.Values{}
	}
	for k, v := range query {
		c.query.Set(k, v)
	}
	return c
}

func (c *HttpClient) Get(path string) *httptest.ResponseRecorder {
	return c.Request("GET", path, nil)
}

func (c *HttpClient) PostJson(path string, data map[string]interface{}) *httptest.ResponseRecorder {
	jsonBytes, _ := json.Marshal(data)
	jsonReader := bytes.NewReader(jsonBytes)

	return c.WithHeader(map[string]string{
		"Content-Type":   "application/json",
		"Content-Length": strconv.Itoa(jsonReader.Len()),
	}).Request("POST", path, jsonReader)
}
func (c *HttpClient) PostForm(path string, data map[string]string) *httptest.ResponseRecorder {
	// https://blog.csdn.net/jeffrey11223/article/details/79819643
	urlData := url.Values{}
	for k, v := range data {
		urlData.Add(k, v)
	}
	urlReader := strings.NewReader(urlData.Encode())

	return c.WithHeader(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}).Request("POST", path, urlReader)
}

func (c *HttpClient) PostFormData(path string, data map[string]string, files map[string]string) *httptest.ResponseRecorder {
	// https://blog.csdn.net/alexanyang/article/details/110200623
	bytesReader := new(bytes.Buffer)
	writer := multipart.NewWriter(bytesReader)
	for k, v := range data {
		err := writer.WriteField(k, v)
		if err != nil {
			panic(err)
		}
	}
	for k, v := range files {
		f, err := os.ReadFile(v)
		if err != nil {
			panic(err)
		}
		fw, err := writer.CreateFormFile(k, filepath.Base(v))
		if err != nil {
			panic(err)
		}
		fw.Write(f)
	}
	writer.Close()

	return c.WithHeader(map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}).Request("POST", path, bytesReader)
}

func (c *HttpClient) PostBinary(path string, body io.Reader) *httptest.ResponseRecorder {
	// Content-Type: image/jpeg
	// Content-Type: application/zip
	return c.Request("POST", path, body)
}

func NewClient(router *gin.Engine) *HttpClient {
	return &HttpClient{Router: router}
}

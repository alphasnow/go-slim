package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-slim/pkg/xhttplog"
	"go-slim/pkg/xlog"
	"io/ioutil"
	"strings"
	"time"
)

const StartTime = "_pkg/xhttplog/start_time"

func NewHttpLog(logManager *xlog.Manager) gin.HandlerFunc {
	zLog := logManager.Logger(xlog.HTTP)
	return func(c *gin.Context) {

		if strings.HasPrefix(c.Request.URL.Path, "/api") == false {
			c.Next()
			return
		}

		// record time
		c.Set(StartTime, time.Now())
		// record response
		w := &xhttplog.ResponseBodyWriter{Body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// url
		msg := fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.String())

		// request
		req := recordRequest(c)
		if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
			req["body"] = limitRecordJsonField(req["body"], 200)
		}
		reqInfo := "Request\n" + fmt.Sprintf("%s\n%s\n%s", req["header"], req["body"], req["extra"])

		c.Next()

		// response
		res := recordResponse(c, w)
		if strings.Contains(c.Writer.Header().Get("Content-Type"), "application/json") {
			res["body"] = limitRecordJsonField(res["body"], 200)
		}
		resInfo := "Response\n" + fmt.Sprintf("%s\n%s\n%s", res["header"], res["body"], res["extra"])

		zLog.Info(msg + "\n" + reqInfo + "\n" + resInfo)
	}
}

func limitRecordJsonField(jsonBody string, limit int) string {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonBody), &data); err != nil {
		return jsonBody
	}

	lim := map[string]interface{}{}
	for k, v := range data {
		if vs, ok := v.(string); ok == true {
			if len(vs) > limit {
				lim[k] = vs[:limit-3] + "..."
				continue
			}
		}
		lim[k] = v
	}
	limJson, _ := json.Marshal(lim)
	return string(limJson)
}

func recordRequest(c *gin.Context) map[string]string {

	headers, _ := json.Marshal(c.Request.Header)

	bodyStr := ""
	if strings.Contains(c.ContentType(), "application/json") {
		bodyIt, _ := readRequestBody(c)
		bodyStr = string(bodyIt.([]byte))
	}

	extra, _ := json.Marshal(map[string]string{
		// "id": requestid.Get(c),
		"ip": c.ClientIP(),
	})
	req := map[string]string{
		"header": string(headers),
		"body":   bodyStr,
		"extra":  string(extra),
	}
	return req

}

func readRequestBody(c *gin.Context) (interface{}, error) {
	var body []byte
	if cb, ok := c.Get(gin.BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
	}
	if body == nil && c.Request.Body != nil {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return nil, err
		}
		c.Set(gin.BodyBytesKey, body)
	}
	return body, nil
}

func recordResponse(c *gin.Context, w *xhttplog.ResponseBodyWriter) map[string]string {
	end := time.Now()

	start, _ := c.Get(StartTime)
	headers, _ := json.Marshal(c.Writer.Header())
	body := w.Body.String()
	extra, _ := json.Marshal(map[string]string{
		"latency": end.Sub(start.(time.Time)).String(),
	})
	resp := map[string]string{
		"header": string(headers),
		"body":   body,
		"extra":  string(extra),
	}
	return resp

}

package xhttplog

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

// https://github.com/gin-gonic/gin/issues/1363
type ResponseBodyWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (r ResponseBodyWriter) Write(b []byte) (int, error) {
	r.Body.Write(b)
	return r.ResponseWriter.Write(b)
}

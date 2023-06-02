package web

import (
	"github.com/stretchr/testify/assert"
	"go-slim/test/common/boot"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Static_Assets(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/storage/assets/normalize.css", nil)
	boot.ServeHTTP(w, req)
	bd := w.Body.String()
	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, bd)
}

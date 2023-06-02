package common

import (
	"encoding/base64"
	"go-slim/internal/app"
	"io/ioutil"
	"path/filepath"
)

func Stubs(f string) string {
	return filepath.Join(app.CallRootPath(), "test/stubs", f)
}
func StubsReadBase64(f string) string {
	imgPath := Stubs(f)
	imgData, err := ioutil.ReadFile(imgPath)
	if err != nil {
		panic(err)
	}
	imgBase64 := base64.StdEncoding.EncodeToString(imgData)
	return imgBase64
}
func StubsWriteBase64(f string, data string) error {
	mask, _ := base64.StdEncoding.DecodeString(data)
	o := Stubs(f)
	err := ioutil.WriteFile(o, mask, 0666)
	return err
}

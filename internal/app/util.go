package app

import (
	"os"
	"path/filepath"
	"runtime"
)

func CallRootPath() string {
	_, file, _, ok := runtime.Caller(0)
	if ok != true {
		panic("Runtime caller error")
	}
	return filepath.Join(filepath.Dir(file), "../../")
}

func ExecRootPath() string {
	binary, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(binary)
}

package cron

import (
	"fmt"
	"go-slim/internal/app"
	"go-slim/pkg/xlog"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// CleanUploadTmpFile
type CleanTmpFileTask struct {
	Log    *zap.Logger
	Path   string
	expire time.Duration
}

func NewCleanTmpFileTask(logManager *xlog.Manager, g *app.Config) *CleanTmpFileTask {
	zLog := logManager.Logger(xlog.CRON)
	tempPath := g.TmpPath
	expire := 30 * time.Minute
	return &CleanTmpFileTask{zLog, tempPath, expire}
}

func (c *CleanTmpFileTask) Handle() {

	files, err := ioutil.ReadDir(c.Path)
	if err != nil {
		return
	}
	expire := time.Now().Add(-c.expire)
	total := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name() == ".gitignore" {
			continue
		}
		if file.ModTime().Before(expire) {
			fi := filepath.Join(c.Path, file.Name())
			err = os.Remove(fi)
			if err != nil {
				c.Log.Error(fmt.Sprintf("CleanUploadTmp Error:%s", fi))
				continue
			}
			total++
		}
	}

	c.Log.Info(fmt.Sprintf("CleanUploadTmp Total:%d", total))

}

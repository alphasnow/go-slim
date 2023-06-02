package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/h2non/filetype"
	"go-slim/internal/app"
	"go-slim/internal/constants"
	"go-slim/internal/models"
	"go-slim/pkg/xsnowflake"
	"go-slim/pkg/xutils"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

type FormFile struct {
	Global    *app.Config
	SnowFlake *xsnowflake.SnowFlake
}

func (s *FormFile) NewModel(fh *multipart.FileHeader, storeDir string) *models.AdminUpload {
	name, _ := s.SnowFlake.Generate()

	tp := s.fileExtension(fh)
	if tp == "" {
		tp = s.fileType(fh)
	}

	file := fmt.Sprintf("%d.%s", name, tp)
	dir := s.dateDir(storeDir)
	up := &models.AdminUpload{
		Disk:       constants.UploadDiskLocal,
		Path:       dir + "/" + file,
		Name:       fh.Filename,
		Type:       tp,
		Size:       uint(fh.Size),
		IsComplete: 0,
	}
	return up
}

func (s *FormFile) LocalUploadDir(storePath string) (string, error) {
	localPath := s.Global.JoinUploads(storePath)

	localDir := filepath.Dir(localPath)
	if err := xutils.MakeDir(localDir); err != nil {
		return "", err
	}
	return localPath, nil
}

func (s *FormFile) fileMd5(data []byte) string {
	ctx := md5.New()
	ctx.Write(data)
	return hex.EncodeToString(ctx.Sum(nil))
}

func (s *FormFile) dateDir(storeDir string) string {
	dateDir := time.Now().Format("20060102")
	return storeDir + "/" + dateDir
}

func (s *FormFile) fileType(fh *multipart.FileHeader) string {
	ext := filepath.Ext(fh.Filename)
	if ext != "" {
		return strings.TrimLeft(ext, ".")
	}
	return ""
}

func (ctl *FormFile) fileExtension(reqFile *multipart.FileHeader) string {
	fp, _ := reqFile.Open()
	defer fp.Close()

	fh := make([]byte, 261)
	_, _ = fp.Read(fh)
	fp.Seek(0, io.SeekStart)

	k, _ := filetype.Match(fh)
	if k != filetype.Unknown {
		return k.Extension
	}
	return ""
}

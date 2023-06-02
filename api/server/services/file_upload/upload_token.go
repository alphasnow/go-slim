package file_upload

import (
	"errors"
	"go-slim/internal/models"
)

type UploadToken struct {
	*UploadInfo `json:",inline,omitempty"`
	ExpireTime  int64 `json:"expire_time"`
}

func (s *UploadToken) Validation(info *UploadInfo) error {
	if s.UploadInfo.FileInfo.FileSize != info.FileSize {
		return errors.New("file size not the same")
	}
	if s.UploadInfo.FileInfo.FileType != info.FileType {
		return errors.New("file type not the same")
	}
	return nil
}

func (s *UploadToken) ValidationByModal(info *models.UserUpload) error {
	uploadInfo := new(UploadInfo)
	uploadInfo.FileSize = info.Size
	uploadInfo.FileType = info.Type
	if err := s.Validation(uploadInfo); err != nil {
		return err
	}
	return nil
}

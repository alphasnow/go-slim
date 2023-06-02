package file_upload

import (
	"encoding/base64"
	"errors"
	"github.com/h2non/filetype"
	"github.com/samber/lo"
)

type UploadInfo struct {
	*FileInfo  `json:",inline,omitempty"`
	*ImageInfo `json:",inline,omitempty"`
}
type FileInfo struct {
	FileName string `json:"file_name" binding:"required"` // image.jpg
	FileSize uint   `json:"file_size" binding:"required"` // 10251
	FileType string `json:"file_type" binding:"required"` // jpg
	FileHash string `json:"file_hash,omitempty"`          // md5
	FileHead string `json:"file_head,omitempty"`          //
}

type ImageInfo struct {
	ImageWidth  uint `json:"image_width,omitempty"`
	ImageHeight uint `json:"image_height,omitempty"`
	ImageDPI    uint `json:"image_dpi,omitempty"`
}

type UploadRule struct {
	MaxSize       uint     `json:"max_size"`
	MaxBase64Size uint     `json:"max_base64_size,omitempty"`
	ValidTypes    []string `json:"valid_types"`
	*ImageRule    `json:",inline,omitempty"`
}
type ImageRule struct {
	MaxWidth  uint `json:"max_width,omitempty"`
	MaxHeight uint `json:"max_height,omitempty"`
	MinWidth  uint `json:"min_width,omitempty"`
	MinHeight uint `json:"min_height,omitempty"`
	MaxLength uint `json:"max_length,omitempty"`
	MinLength uint `json:"min_length,omitempty"`
	MaxRadio  uint `json:"max_radio,omitempty"`
}

func (f *FileInfo) CorrectType() error {
	if f.FileHead == "" {
		return nil
	}

	fh, err := base64.StdEncoding.DecodeString(f.FileHead)
	if err != nil {
		return err
	}

	k, err := filetype.Match(fh)
	if err != nil {
		return err
	}
	if k != filetype.Unknown {
		f.FileType = k.Extension
	}

	return nil
}

func (f *FileInfo) IsImage() bool {
	imgTypes := []string{"jpg", "jp2", "png", "gif", "webp", "tif", "bmp", "psd", "heif", "dwg", "ico"}
	return lo.Contains[string](imgTypes, f.FileType)
}

func (rule *UploadRule) CheckSize(info *FileInfo) error {
	if info.FileSize == 0 {
		return errors.New("upload size is 0")
	}
	if rule.MaxSize == 0 && rule.MaxBase64Size > 0 {
		rule.MaxSize = uint(float32(rule.MaxBase64Size) / 1.333)
	}
	if info.FileSize > rule.MaxSize {
		return errors.New("upload size beyond the limit")
	}

	return nil
}

func (rule *UploadRule) CheckType(info *FileInfo) error {
	if rule.ValidTypes == nil || len(rule.ValidTypes) == 0 {
		return nil
	}

	invalid := true
	for _, v := range rule.ValidTypes {
		if v == info.FileType {
			invalid = false
			break
		}
	}
	if invalid == true {
		return errors.New("upload type invalid")
	}

	return nil
}

func (rule *UploadRule) CheckImage(body *ImageInfo) error {
	if err := rule.checkImageDimension(body); err != nil {
		return err
	}
	if err := rule.checkImageLength(body); err != nil {
		return err
	}
	if err := rule.checkImageRadio(body); err != nil {
		return err
	}
	return nil
}
func (rule *UploadRule) checkImageDimension(body *ImageInfo) error {
	if (rule.MaxWidth > 0 && body.ImageWidth > rule.MaxWidth) || body.ImageWidth < rule.MinWidth {
		return errors.New("upload image width out of limit")
	}
	if (rule.MaxHeight > 0 && body.ImageHeight > rule.MaxHeight) || body.ImageHeight < rule.MinHeight {
		return errors.New("upload image height out of limit")
	}
	return nil
}
func (rule *UploadRule) checkImageLength(body *ImageInfo) error {
	if rule.MaxLength > 0 {
		if body.ImageWidth > rule.MaxLength {
			return errors.New("upload image length out of limit")
		}
		if body.ImageHeight > rule.MaxLength {
			return errors.New("upload image length out of limit")
		}
	}

	if rule.MinLength > 0 {
		if body.ImageWidth < rule.MinLength {
			return errors.New("upload image length out of limit")
		}
		if body.ImageHeight < rule.MinLength {
			return errors.New("upload image length out of limit")
		}
	}
	return nil
}
func (rule *UploadRule) checkImageRadio(body *ImageInfo) error {
	if rule.MaxRadio == 0 {
		return nil
	}
	if body.ImageWidth == 0 || body.ImageHeight == 0 {
		return errors.New("upload image with or height is 0")
	}

	if body.ImageWidth > body.ImageHeight && (body.ImageWidth/body.ImageHeight) > rule.MaxRadio {
		return errors.New("upload image radio out of limit")
	}
	if body.ImageWidth < body.ImageHeight && (body.ImageWidth/body.ImageHeight) < rule.MaxRadio {
		return errors.New("upload image radio out of limit")
	}
	return nil
}

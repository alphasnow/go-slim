package ximage

import (
	"bytes"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
)

type Ext string

const (
	JPEG Ext = "jpeg"
	PNG  Ext = "png"
	GIF  Ext = "gif"
)

func Convert(src []byte, ext Ext) ([]byte, error) {
	srcReader := bytes.NewReader(src)
	srcImg, format, err := image.Decode(srcReader)
	if err != nil {
		return nil, err
	}
	if format == string(ext) {
		return src, nil
	}

	var rstWriter bytes.Buffer
	switch ext {
	case GIF:
		err = gif.Encode(&rstWriter, srcImg, nil)
	case JPEG:
		err = jpeg.Encode(&rstWriter, srcImg, nil)
	case PNG:
		err = png.Encode(&rstWriter, srcImg)
	}
	if err != nil {
		return nil, err
	}
	return rstWriter.Bytes(), nil
}

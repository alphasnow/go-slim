package ximage

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
)

// ByteMaskComposite
func ByteMaskComposite(src []byte, mask []byte) ([]byte, error) {
	maskReader := bytes.NewReader(mask)
	maskImg, _, err := image.Decode(maskReader)
	if err != nil {
		return nil, err
	}

	srcReader := bytes.NewReader(src)
	srcImg, _, err := image.Decode(srcReader)
	if err != nil {
		return nil, err
	}

	rstImg := MaskComposite(srcImg, maskImg.(*image.Gray))
	var rstWriter bytes.Buffer
	err = png.Encode(&rstWriter, rstImg)
	if err != nil {
		return nil, err
	}

	return rstWriter.Bytes(), nil
}

func MaskComposite(src image.Image, mask *image.Gray) *image.RGBA {
	maskAlpha := &GrayAlpha{mask}
	res := image.NewRGBA(src.Bounds())
	draw.DrawMask(res, res.Bounds(), src, image.Point{}, maskAlpha, image.Point{}, draw.Over)
	return res
}

func MaskExport(src *image.NRGBA) image.Image {
	res := &AlphaGray{src}
	return res
}

// GrayAlpha
type GrayAlpha struct {
	IM *image.Gray
}

func (g *GrayAlpha) ColorModel() color.Model {
	return color.AlphaModel
}
func (g *GrayAlpha) Bounds() image.Rectangle {
	return g.IM.Rect
}
func (g *GrayAlpha) At(x, y int) color.Color {
	at := g.IM.GrayAt(x, y)
	return color.Alpha{A: at.Y}
}

// AlphaGray
type AlphaGray struct {
	IM *image.NRGBA
}

func (a *AlphaGray) ColorModel() color.Model {
	return color.GrayModel
}
func (a *AlphaGray) Bounds() image.Rectangle {
	return a.IM.Rect
}
func (a *AlphaGray) At(x, y int) color.Color {
	at := a.IM.NRGBAAt(x, y)
	return color.Gray{Y: at.A}
}

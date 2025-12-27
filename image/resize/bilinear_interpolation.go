package resize

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

type BilinearInterpolation struct {
	Strategy
}

func NewBilinearInterpolation() BilinearInterpolation {
	return BilinearInterpolation{Strategy{}}
}

func (BilinearInterpolation) calculateInterpolation(a, b, c, d, x_weight, y_weight uint32) uint32 {
	interpolation := a*(1-x_weight)*(1-y_weight) + b*x_weight*(1-y_weight) + c*y_weight*(1-x_weight) + d*x_weight*y_weight
	return interpolation
}

func (bi BilinearInterpolation) Resize(im image.Image, width, height int) (image.Image, error) {
	imWidth := float64(im.Bounds().Dx())
	imHeight := float64(im.Bounds().Dy())

	widthf, heightf := bi.calculateNewDimensions(imWidth, imHeight, float64(width), float64(height))

	if widthf == 0 || heightf == 0 {
		return nil, fmt.Errorf("invalid dimension %dx%d", width, height)
	}

	height = int(heightf)
	width = int(widthf)
	widthScaleFactor := (imWidth - 1) / (widthf - 1)
	heightScaleFactor := (imHeight - 1) / (heightf - 1)

	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := range height {
		for x := range width {
			xf, yf := float64(x), float64(y)
			x_l, y_l := int(math.Floor(widthScaleFactor*xf)), int(math.Floor(heightScaleFactor*yf))
			x_h, y_h := int(math.Ceil(widthScaleFactor*xf)), int(math.Ceil(heightScaleFactor*yf))

			xWeight := uint32(int(widthScaleFactor*xf) - x_l)
			yWeight := uint32(int(heightScaleFactor*yf) - y_l)

			ar, ag, ab, aa := im.At(x_l, y_l).RGBA()
			br, bg, bb, ba := im.At(x_h, y_l).RGBA()
			cr, cg, cb, ca := im.At(x_l, y_h).RGBA()
			dr, dg, db, da := im.At(x_h, y_h).RGBA()

			pixelR := uint16(bi.calculateInterpolation(ar, br, cr, dr, xWeight, yWeight))
			pixelG := uint16(bi.calculateInterpolation(ag, bg, cg, dg, xWeight, yWeight))
			pixelB := uint16(bi.calculateInterpolation(ab, bb, cb, db, xWeight, yWeight))
			pixelA := uint16(bi.calculateInterpolation(aa, ba, ca, da, xWeight, yWeight))

			newImage.Set(x, y, color.RGBA64{pixelR, pixelG, pixelB, pixelA})
		}
	}

	return newImage, nil
}

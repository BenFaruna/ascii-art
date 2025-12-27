package resize

import (
	"fmt"
	"image"
)

type NearestNeighbor struct {
	Strategy
}

func NewNearestNeighbor() NearestNeighbor {
	return NearestNeighbor{Strategy{}}
}

func (n NearestNeighbor) Resize(im image.Image, width, height int) (image.Image, error) {
	imWidth := float64(im.Bounds().Dx())
	imHeight := float64(im.Bounds().Dy())

	widthf, heightf := n.calculateNewDimensions(imWidth, imHeight, float64(width), float64(height))
	width = int(widthf)
	height = int(heightf)

	if width == 0 && height == 0 {
		return nil, fmt.Errorf("invalid dimension %dx%d", width, height)
	}

	widthScaleFactor := (imWidth - 1) / (widthf - 1)
	heightScaleFactor := (imHeight - 1) / (heightf - 1)

	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := range height {
		for x := range width {
			srcX := int(float64(x) * widthScaleFactor)
			srcY := int(float64(y) * heightScaleFactor)

			newImage.Set(x, y, im.At(srcX, srcY))
		}
	}

	return newImage, nil
}

package resize

import (
	"image"
)

type ResizeStrategy interface {
	Resize(im image.Image, width, height int) (image.Image, error)
}

type Strategy struct{}

func (s Strategy) calculateAspectRatio(left, right float64) float64 {
	return left / right
}

func (s Strategy) calculateNewDimensions(originalWidth, originalHeight, targetWidth, targetHeight float64) (float64, float64) {
	if targetWidth == 0 && targetHeight == 0 {
		return originalWidth, originalHeight
	}

	if targetWidth == 0 {
		targetWidth = targetHeight * s.calculateAspectRatio(originalWidth, originalHeight)
	} else if targetHeight == 0 {
		targetHeight = targetWidth * s.calculateAspectRatio(originalHeight, originalWidth)
	}

	return targetWidth, targetHeight
}

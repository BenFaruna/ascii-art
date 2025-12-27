package resize

import (
	"image"
)

type Resizer interface {
	ResizeImage(width, height int) (image.Image, error)
	SetStrategy(s ResizeStrategy)
}

type ResizerStrategy struct {
	im       image.Image
	Strategy ResizeStrategy
}

func NewResizerStrategy(image image.Image, strategy ResizeStrategy) *ResizerStrategy {
	return &ResizerStrategy{
		im:       image,
		Strategy: strategy,
	}
}

func (r *ResizerStrategy) ResizeImage(width, height int) (image.Image, error) {
	return r.Strategy.Resize(r.im, width, height)
}

func (r *ResizerStrategy) SetStrategy(s ResizeStrategy) {
	r.Strategy = s
}

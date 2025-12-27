package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"time"

	"github.com/benfaruna/ascii-art/image/resize"
)

func selectResizer(res string) resize.ResizeStrategy {
	switch res {
	case "low":
		return resize.NewNearestNeighbor()
	case "high":
		return resize.NewBilinearInterpolation()
	default:
		return resize.NewNearestNeighbor()
	}
}

func saveImage(im image.Image, path string, format string) error {
	if path == "" {
		path = fmt.Sprintf("%d.%s", time.Now().Unix(), format)
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case "png":
		return png.Encode(file, im)
	case "jpeg":
		return jpeg.Encode(file, im, &jpeg.Options{Quality: 90})
	case "jpg":
		return jpeg.Encode(file, im, &jpeg.Options{Quality: 90})
	default:
		return errors.New("unsupported format")
	}
}

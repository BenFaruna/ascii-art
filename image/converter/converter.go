package converter

import (
	"image"
	color_lib "image/color"
)

func ConvertToGrayscale(img image.Image) image.Image {
	// Implementation of ConvertToGrayscale function
	// Create a new grayscale image with the same dimensions as the input image
	grayImg := image.NewGray16(img.Bounds())

	// Iterate over each pixel in the input image
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			// Get the color of the current pixel
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()

			// Convert the color to grayscale using the formula: 0.299*R + 0.587*G + 0.114*B
			gray := uint16(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))

			// Set the grayscale value for the current pixel in the grayscale image
			grayImg.Set(x, y, color_lib.Gray16{gray})
		}
	}

	return grayImg
}

func ConvertToAscii(img image.Image) string {
	grayRamp := "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'."
	var rampLength int = len(grayRamp)

	getAsciiChar := func(gray int) string {
		val := grayRamp[(rampLength-1)*gray/255%rampLength]
		return string(val)
	}

	// Convert the grayscale image to ASCII art
	var asciiArt string
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			gray := img.(*image.Gray16).Gray16At(x, y).Y
			asciiArt += getAsciiChar(int(gray))
		}
		asciiArt += "\n"
	}

	return asciiArt
}

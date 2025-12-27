package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/benfaruna/ascii-art/image/converter"
)

var (
	filepath   = flag.String("filepath", "", "The file path to process")
	width      = flag.Int("width", 0, "resize width (optional)")
	height     = flag.Int("height", 0, "resize height (optional)")
	resolution = flag.String("res", "low", "resolution (low, high)")
	output     = flag.String("output", "", "output file path (optional)")
	outputType = flag.String("outputType", "ascii", "output type (ascii, grayscale, color)")
)

func main() {
	flag.Parse()

	if *filepath == "" {
		flag.PrintDefaults()
		return
	}

	file, err := os.Open(*filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	im, format, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	if *width > 0 || *height > 0 {
		resizer := selectResizer(*resolution)
		im, err = resizer.Resize(im, *width, *height)
		if err != nil {
			fmt.Println("Error resizing image:", err)
			return
		}
	}

	var asciiOutput string
	switch *outputType {
	case "color":
	case "grayscale":
		im = converter.ConvertToGrayscale(im)
	case "ascii":
		im = converter.ConvertToGrayscale(im)
		asciiOutput = converter.ConvertToAscii(im)
	default:
		fmt.Println("Invalid output type:", *outputType)
		return
	}

	if *outputType == "ascii" {
		if *output == "" {
			fmt.Println(asciiOutput)
			return
		}
		file, err := os.Create(*output)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer file.Close()
		fmt.Fprintln(file, asciiOutput)
		return
	}

	err = saveImage(im, *output, format)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}
}

package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/benfaruna/ascii-art/image/converter"
	"github.com/benfaruna/ascii-art/image/resize"
)

var (
	filename = flag.String("filename", "", "The name of the file to process")
)

func main() {
	flag.Parse()

	if *filename == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	im, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		os.Exit(1)
	}

	nn := resize.NewNearestNeighbor()
	bi := resize.NewBilinearInterpolation()

	nnImage, _ := nn.Resize(im, 500, 500)
	biImage, _ := bi.Resize(im, 500, 500)

	nnFile, _ := os.OpenFile(fmt.Sprintf("%s_nn.jpeg", *filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	jpeg.Encode(nnFile, nnImage, &jpeg.Options{Quality: 90})
	nnFile.Sync()
	nnFile.Close()

	biFile, _ := os.OpenFile(fmt.Sprintf("%s_bi.jpeg", *filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	jpeg.Encode(biFile, biImage, &jpeg.Options{Quality: 90})
	biFile.Sync()
	biFile.Close()

	gray_img := converter.ConvertToGrayscale(biImage)
	file, _ = os.OpenFile("gray.jpeg", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	jpeg.Encode(file, gray_img, &jpeg.Options{Quality: 90})
	file.Sync()
	file.Close()

	art := converter.ConvertToAscii(gray_img)
	file, _ = os.OpenFile(fmt.Sprintf("%s.txt", *filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	file.WriteString(art)
	file.Sync()
	file.Close()
}

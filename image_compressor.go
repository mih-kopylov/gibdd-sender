package main

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
)

func resizeImage(imageFileName string) io.Reader {
	imageFile, err := os.Open(imageFileName)
	if err != nil {
		log.Panic(err)
	}
	defer imageFile.Close()
	decodedImage, _, err := image.Decode(imageFile)
	if err != nil {
		log.Panic(err)
	}
	resizedImage := resize.Resize(2000, 0, decodedImage, resize.Lanczos3)
	var buffer bytes.Buffer
	err = jpeg.Encode(&buffer, resizedImage, nil)
	if err != nil {
		log.Panic(err)
	}
	return &buffer
}

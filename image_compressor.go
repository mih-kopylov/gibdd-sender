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

func resizeImage(imageFile *os.File) io.Reader {
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

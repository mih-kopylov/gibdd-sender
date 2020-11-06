package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// returns array of images in current directory
func getImages(imagesDirectory string) []string {
	var result []string
	log.Println("searching for images in directory", imagesDirectory)
	fileInfos, err := ioutil.ReadDir(imagesDirectory)
	if err != nil {
		log.Panic(err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		ext := filepath.Ext(fileInfo.Name())
		if !strings.EqualFold(ext, ".jpg") && !strings.EqualFold(ext, ".jpeg") {
			continue
		}
		result = append(result, fileInfo.Name())
		log.Println("found image", fileInfo.Name())
	}
	if len(result) == 0 {
		log.Fatal("No files found")
	}
	return result
}

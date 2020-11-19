package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func createArchive(imagesDirectory string, images []string) string {
	archiveName := filepath.Base(imagesDirectory) + ".zip"
	log.Println("create archive", archiveName)
	file, err := os.Create(filepath.Join(imagesDirectory, archiveName))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	for i, imageFileName := range images {
		imageFile, err := os.Open(filepath.Join(imagesDirectory, imageFileName))
		if err != nil {
			log.Panic(err)
		}
		zipImageWriter, _ := zipWriter.Create(imageFileName)
		log.Println("adding file to archive", imageFileName)
		_, err = io.Copy(zipImageWriter, resizeImage(imageFile))
		if err != nil {
			log.Panic(err)
		}
		log.Println("file added", imageFileName, (i+1)*100/len(images), "%")
	}
	return archiveName
}

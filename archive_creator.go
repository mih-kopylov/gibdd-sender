package main

import (
	"archive/zip"
	"github.com/signintech/gopdf"
	"io"
	"log"
	"os"
	"path/filepath"
)

func createZipArchive(imagesDirectory string, images []string) string {
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
		zipImageWriter, _ := zipWriter.Create(imageFileName)
		log.Println("adding file to archive", imageFileName)
		_, err = io.Copy(zipImageWriter, resizeImage(filepath.Join(imagesDirectory, imageFileName)))
		if err != nil {
			log.Panic(err)
		}
		log.Println("file added", imageFileName, (i+1)*100/len(images), "%")
	}
	return archiveName
}

func createPdfArchive(imagesDirectory string, images []string) string {
	archiveName := filepath.Base(imagesDirectory) + ".pdf"
	log.Println("create archive", archiveName)

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 1170, H: 1170}})
	for i, imageFileName := range images {
		log.Println("adding file to archive", imageFileName)
		pdf.AddPage()
		imageHolder, err := gopdf.ImageHolderByReader(resizeImage(filepath.Join(imagesDirectory, imageFileName)))
		err = pdf.ImageByHolder(imageHolder, 20, 20, nil)
		if err != nil {
			log.Panic(err)
		}
		log.Println("file added", imageFileName, (i+1)*100/len(images), "%")
	}

	err := pdf.WritePdf(filepath.Join(imagesDirectory, archiveName))
	if err != nil {
		log.Panic(err)
	}
	return archiveName
}

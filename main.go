package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	imagesDirectory := os.Args[1]
	configFile := os.Args[2]
	configuration := readConfiguration(configFile)
	images := getImages(imagesDirectory)
	var archiveFile string
	switch archiveType := strings.ToLower(configuration.ArchiveType); archiveType {
	case "zip":
		archiveFile = createZipArchive(imagesDirectory, images)
	case "pdf":
		archiveFile = createPdfArchive(imagesDirectory, images)
	default:
		log.Fatal("Unsupported archive type: ", archiveType)
	}
	timeAddress := getTimeAddress(imagesDirectory)
	switch rec := strings.ToLower(configuration.Receiver); rec {
	case "mvd":
		log.Println("sending a message to MVD receiver")
		sendMessageToMvd(imagesDirectory, configuration, timeAddress, archiveFile)
	default:
		log.Fatal("Unsupported receiver: ", rec)
	}
}

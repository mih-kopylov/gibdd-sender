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
	archive := createArchive(imagesDirectory, images)
	timeAddress := getTimeAddress(imagesDirectory)
	switch rec := strings.ToLower(configuration.Receiver); rec {
	case "mvd":
		log.Println("sending a message to MVD receiver")
		sendMessageToMvd(imagesDirectory, configuration, timeAddress, archive)
	default:
		log.Fatal("Unsupported receiver: ", rec)
	}
}

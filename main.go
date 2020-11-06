package main

import "os"

func main() {
	println("hello")
	imagesDirectory := os.Args[1]
	configFile := os.Args[2]
	images := getImages(imagesDirectory)
	archive := createArchive(imagesDirectory, images)
	timeAddress := getTimeAddress(imagesDirectory)
	sendMessage(imagesDirectory, configFile, timeAddress, archive)
}

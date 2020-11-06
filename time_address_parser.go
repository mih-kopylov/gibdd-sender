package main

import (
	"log"
	"path/filepath"
	"strings"
	"time"
)

func getTimeAddress(directoryName string) TimeAddress {
	directoryNameBase := filepath.Base(directoryName)
	return parseDirectoryName(directoryNameBase)
}

func parseDirectoryName(directoryName string) TimeAddress {
	parts := strings.Split(directoryName, " ")
	directoryTime, err := time.Parse("2006-01-02 15-04", parts[0]+" "+parts[1])
	if err != nil {
		log.Panic(err)
	}
	address := strings.Join(parts[2:], " ")
	return TimeAddress{directoryTime, address}
}

type TimeAddress struct {
	Time    time.Time
	Address string
}

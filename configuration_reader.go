package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func readConfiguration(configFile string) Configuration {
	fileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panic(err)
	}
	var result Configuration
	err = json.Unmarshal(fileContent, &result)
	if err != nil {
		log.Panic(err)
	}
	return result
}

type Configuration struct {
	LastName        string `json:"lastName"`
	Name            string `json:"name"`
	MiddleName      string `json:"middleName"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	RegionId        string `json:"regionId"`
	Subunit         string `json:"subunit"`
	MessageTemplate string `json:"messageTemplate"`
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"path/filepath"
)

func sendMessage(imagesDirectory string, configuratoin Configuration, timeAddress TimeAddress, archiveFileName string) {
	log.Println("sending request with archive", archiveFileName)

	request := createRequest(imagesDirectory, configuratoin, timeAddress, archiveFileName)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	responseBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	var response ResponseDto
	err = json.Unmarshal(responseBodyBytes, &response)
	if err != nil {
		dumpRequest, _ := httputil.DumpRequest(request, true)
		log.Println(string(dumpRequest))
		log.Println(string(responseBodyBytes))
		log.Panic(err)
	}
	if response.ErrorCode != 0 {
		dumpRequest, _ := httputil.DumpRequest(request, true)
		log.Println(string(dumpRequest))
		log.Fatal(fmt.Sprintf("%d: %s", response.ErrorCode, response.ErrorName))

	} else {
		log.Println(response.ErrorName)
	}
}

func createRequest(imagesDirectory string, configuration Configuration, timeAddress TimeAddress, archiveFileName string) *http.Request {
	message := formatMessage(configuration.MessageTemplate, timeAddress)

	payload := new(bytes.Buffer)
	payloadWriter := multipart.NewWriter(payload)
	filePart, err := payloadWriter.CreateFormFile("file", archiveFileName)
	if err != nil {
		log.Panic(err)
	}
	fileContent, err := ioutil.ReadFile(filepath.Join(imagesDirectory, archiveFileName))
	if err != nil {
		log.Panic(err)
	}
	_, err = filePart.Write(fileContent)
	if err != nil {
		log.Panic(err)
	}
	_ = payloadWriter.WriteField("declarant[f]", configuration.LastName)
	_ = payloadWriter.WriteField("declarant[i]", configuration.Name)
	_ = payloadWriter.WriteField("declarant[o]", configuration.MiddleName)
	_ = payloadWriter.WriteField("answer_addr[email]", configuration.Email)
	_ = payloadWriter.WriteField("version", "2")
	_ = payloadWriter.WriteField("phone", configuration.Phone)
	_ = payloadWriter.WriteField("region_id", configuration.RegionId)
	_ = payloadWriter.WriteField("address[subunit]", configuration.Subunit)
	_ = payloadWriter.WriteField("message", message)
	request, _ := http.NewRequest(http.MethodPost, "https://mvd.ru/api/v2/request", payload)
	request.Header.Add("Content-Type", "multipart/form-data; boundary="+payloadWriter.Boundary())
	//these headers below are required to emulate official MVD Android application
	request.Header.Add("User-Agent", "android")
	request.Header.Add("Application", "sitesoft")
	return request
}

func formatMessage(messageTemplate string, timeaddress TimeAddress) string {
	parsedTemplate, err := template.New("messageTemplate").Parse(messageTemplate)
	if err != nil {
		log.Panic(err)
	}
	buffer := bytes.Buffer{}
	err = parsedTemplate.Execute(&buffer, timeaddress)
	if err != nil {
		log.Panic(err)
	}
	return buffer.String()
}

type ResponseDto struct {
	ErrorCode int    `json:"error_code"`
	ErrorName string `json:"error_name"`
	Data      Data   `json:"data"`
}

type Data struct {
	Code string `json:"code"`
}

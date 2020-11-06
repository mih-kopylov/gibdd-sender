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

func sendMessage(imagesDirectory string, configFile string, timeAddress TimeAddress, archiveFileName string) {
	log.Println("sending request with archive", archiveFileName)

	request := createRequest(imagesDirectory, configFile, timeAddress, archiveFileName)
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

func createRequest(imagesDirectory string, configFile string, timeAddress TimeAddress, archiveFileName string) *http.Request {
	personalData := readPersonalData(configFile)
	message := formatMessage(personalData.MessageTemplate, timeAddress)

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
	_ = payloadWriter.WriteField("declarant[f]", personalData.LastName)
	_ = payloadWriter.WriteField("declarant[i]", personalData.Name)
	_ = payloadWriter.WriteField("declarant[o]", personalData.MiddleName)
	_ = payloadWriter.WriteField("answer_addr[email]", personalData.Email)
	_ = payloadWriter.WriteField("version", "2")
	_ = payloadWriter.WriteField("phone", personalData.Phone)
	_ = payloadWriter.WriteField("region_id", personalData.RegionId)
	_ = payloadWriter.WriteField("address[subunit]", personalData.Subunit)
	_ = payloadWriter.WriteField("message", message)
	request, _ := http.NewRequest(http.MethodPost, "https://mvd.ru/api/v2/request", payload)
	request.Header.Add("Content-Type", "multipart/form-data; boundary="+payloadWriter.Boundary())
	//these headers below are required to emulate official MVD Android application
	request.Header.Add("User-Agent", "android")
	request.Header.Add("Application", "sitesoft")
	return request
}

func readPersonalData(configFile string) PersonalData {
	fileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panic(err)
	}
	var result PersonalData
	err = json.Unmarshal(fileContent, &result)
	if err != nil {
		log.Panic(err)
	}
	return result
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

type PersonalData struct {
	LastName        string `json:"lastName"`
	Name            string `json:"name"`
	MiddleName      string `json:"middleName"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	RegionId        string `json:"regionId"`
	Subunit         string `json:"subunit"`
	MessageTemplate string `json:"messageTemplate"`
}

type ResponseDto struct {
	ErrorCode int    `json:"error_code"`
	ErrorName string `json:"error_name"`
	Data      Data   `json:"data"`
}

type Data struct {
	Code string `json:"code"`
}

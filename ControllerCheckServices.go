package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var responseDataCheckService = make(map[string]CheckServicesController)

type CheckServicesController struct {
	ID      string `json:"ID"`
	STATE   string `json:"STATE"`
	SUPER   string `json:"SUPER"`
	MESSAGE string `json:"MESSAGE"`
	POINTS  string `json:"POINTS"`
}

func NewCheckServicesControllerEmpty() *CheckServicesController {
	return &CheckServicesController{
		ID:      "0",
		STATE:   "0",
		SUPER:   "0",
		MESSAGE: "0",
		POINTS:  "0",
	}
}

func NewCheckServicesController(id string, state string, super string, message string, points string) *CheckServicesController {
	return &CheckServicesController{
		ID:      id,
		STATE:   state,
		SUPER:   super,
		MESSAGE: message,
		POINTS:  points,
	}
}

func CheckServicesRequestGet(w http.ResponseWriter, r *http.Request) {

	dataToSend := NewCheckServicesControllerEmpty()
	dataToSend.ID = "0"
	dataToSend.MESSAGE = responseDataCheckService["0"].MESSAGE
	dataToSend.POINTS = responseDataCheckService["0"].POINTS
	dataToSend.STATE = responseDataCheckService["0"].STATE
	dataToSend.SUPER = responseDataCheckService["0"].SUPER

	//header
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(dataToSend)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func CheckServicesRequestPost(w http.ResponseWriter, r *http.Request) {
	var data CheckServicesController

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	s := buf.String()
	fmt.Println("---------")
	fmt.Println(s)
	byteArray := []byte(s)
	//key_init := []byte("{") //123
	//key_end := []byte("}")  //125

	position := 0
	sum := 0
	for _, v := range byteArray {
		fmt.Println(v)
		if v == 123 { // 123 es { en byte
			position = sum
		}
		sum += 1
	}

	dataJson := string(byteArray[position:])
	fmt.Println(dataJson)

	fmt.Println("-----end----")

	err := json.Unmarshal([]byte(dataJson), &data)

	if err != nil {
		log.Fatal(err)
		log.Println("json.Compact:", err)
		if serr, ok := err.(*json.SyntaxError); ok {
			log.Println("Occurred at offset:", serr.Offset)
			// … something to show the data in buff around that offset …
		}
	}

	responseDataCheckService["0"] = data

	dataToSend := NewResponseControllerEmpty()
	dataToSend.ID = "0"
	dataToSend.STATE = "1"
	dataToSend.ERROR = "0"
	dataToSend.CODE = "200"

	//header
	w.Header().Set("Content-Type", "application/json")
	//resp, err := json.Marshal(data)
	resp, err := json.Marshal(dataToSend)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

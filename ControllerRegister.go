package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var responseDataRegister = make(map[string]RegisterController)
var idRegister int

type RegisterController struct {
	ID       string
	PHONE    string
	USERNAME string
	PASSWORD string
	CREATED  string
}

func NewRegisterController(id string, phone string, username string, pass string) *RegisterController {
	return &RegisterController{
		ID:       id,
		PHONE:    phone,
		USERNAME: username,
		PASSWORD: pass,
		CREATED:  time.Now().String(),
	}
}

func RegisterRequestGet(w http.ResponseWriter, r *http.Request) {
	var data []RegisterController
	for _, value := range responseDataRegister {
		data = append(data, value)
	}

	//header
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
func RegisterRequestPost(w http.ResponseWriter, r *http.Request) {
	var data RegisterController

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
	idDataGame++
	id := strconv.Itoa(idDataGame)
	fmt.Println(data)
	responseDataRegister[id] = data

	//header
	dataToSend := NewResponseController("0", "1", "0", "200")
	w.Header().Set("Content-Type", "application/json")
	//resp, err := json.Marshal(data)
	resp, err := json.Marshal(dataToSend)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
func RegisterRequestPostAdmin(w http.ResponseWriter, r *http.Request) {
	var data RegisterController
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	idRegister++
	id := strconv.Itoa(idRegister)
	data.CREATED = time.Now().String()
	responseDataRegister[id] = data

	//header
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
func RegisterRequestUpdate(w http.ResponseWriter, r *http.Request) {

	var newData RegisterController
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		panic(err)
	}
	id := newData.ID

	if lastData, ok := responseDataRegister[id]; ok {
		newData.CREATED = lastData.CREATED
		delete(responseDataRegister, id)
		responseDataRegister[id] = newData
	} else {
		log.Printf("no se encontró el id: %s", id)
	}

	//header
	dataToSend := NewResponseController("0", "1", "0", "200")
	w.Header().Set("Content-Type", "application/json")
	//resp, err := json.Marshal(data)
	resp, err := json.Marshal(dataToSend)
	if err != nil {
		panic(err)
	}
	//w.WriteHeader(http.StatusNoContent)
	w.Write(resp)
}

//func RegisterRequestDelete() {}

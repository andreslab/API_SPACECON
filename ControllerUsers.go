package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

var responseDataUsers = make(map[string]UsersController)
var idUsers int

type UsersController struct {
	ID       string `json:"ID"`
	USERNAME string `json:"USERNAME"`
	PHONE    string `json:"PHONE"`
	CREATED  string `json:"CREATED"`
	DIAMONDS string `json:"DIAMONDS"`
	STATE    string `json:"STATE"`
}

func NewUsersController(id string, username string, phone, string, numDiamond string, state string) *UsersController {
	return &UsersController{
		ID:       id,
		USERNAME: username,
		PHONE:    phone,
		CREATED:  time.Now().String(),
		DIAMONDS: "0",
		STATE:    "0",
	}
}

func UsersRequestGet(w http.ResponseWriter, r *http.Request) {
	var data []UsersController
	for _, value := range responseDataUsers {
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
func UsersRequestPost(w http.ResponseWriter, r *http.Request) {
	var data UsersController
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	userExist := false

	for index := range responseDataUsers {
		if responseDataUsers[index].USERNAME == data.USERNAME {
			userExist = true
		}
	}

	dataToSend := NewResponseControllerEmpty()

	if !userExist {

		//save new user

		id := strconv.Itoa(idUsers)
		data.ID = id
		responseDataUsers[id] = data
		idUsers++

		dataToSend.ID = "0"
		dataToSend.STATE = "1"
		dataToSend.ERROR = "0"
		dataToSend.CODE = "200"
	} else {
		dataToSend.ID = "0"
		dataToSend.STATE = "0"
		dataToSend.ERROR = "1"
		dataToSend.CODE = "400"
	}

	//header
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(dataToSend)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}
func UsersRequestUpdate(w http.ResponseWriter, r *http.Request) {
	var newData UsersController
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		panic(err)
	}
	id := newData.ID

	if lastData, ok := responseDataUsers[id]; ok {
		newData.CREATED = lastData.CREATED
		delete(responseDataUsers, id)
		responseDataUsers[id] = newData
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

func UsersRequestPostAdmin(w http.ResponseWriter, r *http.Request) {
	var data UsersController
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	id := strconv.Itoa(idUsers)
	data.CREATED = time.Now().String()
	responseDataUsers[id] = data
	idUsers++

	//header
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func UsersRequestDeleteAdmin(w http.ResponseWriter, r *http.Request) {
	idUsers = 0

	for index := range responseDataUsers {
		delete(responseDataUsers, index)
	}

	dataToSend := NewResponseControllerEmpty()

	if len(responseDataDataGame) < 1 {
		dataToSend.ID = "0"
		dataToSend.STATE = "1"
		dataToSend.ERROR = "0"
		dataToSend.CODE = "200"
	} else {
		dataToSend.ID = "0"
		dataToSend.STATE = "0"
		dataToSend.ERROR = "1"
		dataToSend.CODE = "400"
	}

	//header
	w.Header().Set("Content-Type", "application/json")
	//resp, err := json.Marshal(data)
	resp, err := json.Marshal(dataToSend)
	if err != nil {
		panic(err)
	}
	//w.WriteHeader(http.StatusNoContent)
	w.Write(resp)
}

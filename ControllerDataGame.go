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

var responseDataDataGame = make(map[string]DataGameController)
var idDataGame int

type DataGameController struct {
	ID        string `json:"ID"`
	NAME      string `json:"NAME"`
	VALUE     string `json:"VALUE"`
	LATITUDE  string `json:"LATITUDE"`
	LONGITUDE string `json:"LONGITUDE"`
	STATE     string `json:"STATE"`
	CREATED   string `json:"CREATED"`
	HUNTER    string `json:"HUNTER"`
}

func NewDataGameControllerEmpty() *DataGameController {
	return &DataGameController{
		ID:        "0",
		NAME:      "none",
		VALUE:     "0",
		LATITUDE:  "0",
		LONGITUDE: "0",
		STATE:     "0",
		CREATED:   time.Now().String(),
		HUNTER:    "none",
	}
}

func NewDataGameController(id string, name string, value string, latitude string, longitude string, state string, hunter string) *DataGameController {
	return &DataGameController{
		ID:        id,
		NAME:      name,
		VALUE:     value,
		LATITUDE:  latitude,
		LONGITUDE: longitude,
		STATE:     state,
		CREATED:   time.Now().String(),
		HUNTER:    hunter,
	}
}

func DataGameRequestGet(w http.ResponseWriter, r *http.Request) {
	var data []DataGameController
	for _, value := range responseDataDataGame {
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

/*func DataGameRequestPost(w http.ResponseWriter, r *http.Request) {
	var data DataGameController

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

	id := strconv.Itoa(idDataGame)
	fmt.Println(data)
	responseDataDataGame[id] = data
	idDataGame++

	//header
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}*/

//sin parametros en la url
func DataGameRequestUpdate(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)
	id := vars["id"]*/
	var newData DataGameController
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		panic(err)
	}
	id := newData.ID
	fmt.Println(newData.LATITUDE)

	if lastData, ok := responseDataDataGame[id]; ok {
		newData.CREATED = lastData.CREATED
		delete(responseDataDataGame, id)
		responseDataDataGame[id] = newData
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

//con parametros en la url
func DataGameRequestUpdateAdmin(w http.ResponseWriter, r *http.Request) {
	var newData DataGameController
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		panic(err)
	}
	id := newData.ID

	if lastData, ok := responseDataDataGame[id]; ok {
		newData.CREATED = lastData.CREATED
		delete(responseDataDataGame, id)
		responseDataDataGame[id] = newData
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

func DataGameRequestPostAdmin(w http.ResponseWriter, r *http.Request) {

	var data DataGameController

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

	//save new object

	id := strconv.Itoa(idDataGame)
	data.ID = id
	data.CREATED = time.Now().String()
	fmt.Println(data)
	responseDataDataGame[id] = data
	idDataGame++

	dataToSend := NewResponseControllerEmpty()

	dataToSend.ID = "0"
	dataToSend.STATE = "1"
	dataToSend.ERROR = "0"
	dataToSend.CODE = "200"

	//header
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(dataToSend)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func DataGameRequestDeleteAdmin(w http.ResponseWriter, r *http.Request) {

	for index := range responseDataDataGame {
		delete(responseDataDataGame, index)
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

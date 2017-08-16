package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var responseDataLogin = make(map[string]LoginController)
var idLogin int

type LoginController struct {
	USERNAME string
	PASSWORD string
}

func NewLoginController(username string, password string) *LoginController {
	return &LoginController{
		USERNAME: username,
		PASSWORD: password,
	}
}

//func LoginRequestGet(w http.ResponseWriter, r *http.Request) {}
func LoginRequestPost(w http.ResponseWriter, r *http.Request) {
	var data LoginController

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
	responseDataLogin[id] = data

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

func LoginRequestPostAdmin(w http.ResponseWriter, r *http.Request) {
	var data LoginController
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	idLogin++
	id := strconv.Itoa(idLogin)
	responseDataLogin[id] = data

	//header
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

//func LoginRequestUpdate()    {}
//func LoginRequestDelete() {}

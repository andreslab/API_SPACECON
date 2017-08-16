package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	controller "./controller"
	//servicedata "./resources/database"
)

func CheckServicesRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controller.UsersRequestGet(w, r)
	case "POST":
	case "PUT":
	default:
	}
}

func UsersRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controller.UsersRequestGet(w, r)
	case "POST":
		controller.UsersRequestPost(w, r)
	case "PUT":
	default:
	}
}

func UsersRequestAdmin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controller.UsersRequestGet(w, r)
	case "POST":
		controller.UsersRequestPostAdmin(w, r)
	case "PUT":
	default:
	}
}

func DataGameRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controller.DataGameRequestGet(w, r)
	case "POST":
		controller.DataGameRequestPost(w, r)
	case "PUT":
		controller.DataGameRequestUpdate(w, r)
	default:
	}
}

func DataGameRequestAdmin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controller.DataGameRequestGet(w, r)
	case "POST":
		controller.DataGameRequestPostAdmin(w, r)
	case "PUT":
		//controller.DataGameRequestUpdate(w, r)
	default:
	}
}

func RegisterRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controller.RegisterRequestGet(w, r)
	case "POST":
		controller.RegisterRequestPost(w, r)
	case "PUT":
		controller.RegisterRequestUpdate(w, r)
	default:
	}
}

func RegisterRequestAdmin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controller.RegisterRequestGet(w, r)
	case "POST":
		controller.RegisterRequestPostAdmin(w, r)
	case "PUT":
		//controller.RegisterRequestUpdate(w, r)
	default:
	}
}

func LoginRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
		controller.LoginRequestPost(w, r)
	case "PUT":
	default:
	}
}

func main() {
	fmt.Print("...init....")

	//database:
	/*data := controller.NewDataGameController(
		"0",
		"diamante",
		"0",
		"2.3432",
		"-2.4646",
		"0",
		"none")

	fmt.Println(data.NAME)*/

	//servicedata.CreateTableData()

	/*servicedata.InsertTableData(data)

	d := servicedata.SelectTableData()
	fmt.Println(d.NAME)*/

	//end database

	/*
		// with library gorilla mux
		r := mux.NewRouter().StrictSlash(false)

		r.HandleFunc("/odisea/check", controller.CheckServicesRequestGet).Methods("GET")

		r.HandleFunc("/odisea/register", controller.RegisterRequestGet).Methods("GET")
		r.HandleFunc("/odisea/register", controller.RegisterRequestPost).Methods("POST")
		r.HandleFunc("/odisea/admin/register", controller.RegisterRequestPostAdmin).Methods("POST")
		r.HandleFunc("/odisea/register", controller.RegisterRequestUpdate).Methods("PUT")

		r.HandleFunc("/odisea/login", controller.LoginRequestPost).Methods("POST")

		r.HandleFunc("/odisea/data", controller.DataGameRequestGet).Methods("GET")
		r.HandleFunc("/odisea/data", controller.DataGameRequestPost).Methods("POST")
		r.HandleFunc("/odisea/admin/data", controller.DataGameRequestPostAdmin).Methods("POST")
		r.HandleFunc("/odisea/data", controller.DataGameRequestUpdate).Methods("PUT")

		r.HandleFunc("/odisea/users", controller.UsersRequestGet).Methods("GET")
		r.HandleFunc("/odisea/users", controller.UsersRequestPost).Methods("POST")
		r.HandleFunc("/odisea/admin/users", controller.UsersRequestPostAdmin).Methods("POST")
		r.HandleFunc("/odisea/users", controller.UsersRequestUpdate).Methods("PUT")*/

	mux := http.NewServeMux()

	mux.HandleFunc("/odisea/check", CheckServicesRequest)
	mux.HandleFunc("/odisea/register", RegisterRequest)
	mux.HandleFunc("/odisea/admin/register", RegisterRequestAdmin)
	mux.HandleFunc("/odisea/data", DataGameRequest)
	mux.HandleFunc("/odisea/admin/data", DataGameRequestAdmin)
	mux.HandleFunc("/odisea/login", LoginRequest)
	mux.HandleFunc("/odisea/users", UsersRequest)
	mux.HandleFunc("/odisea/admin/users", UsersRequestAdmin)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	log.Printf("init server ....")
	server.ListenAndServe()
}

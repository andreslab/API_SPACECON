package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	//controller "./controller"
	//servicedata "./resources/database"
)

func CheckServicesRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		CheckServicesRequestGet(w, r)
	case "POST":
	case "PUT":
	default:
	}
}

func UsersRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		UsersRequestGet(w, r)
	case "POST":
		UsersRequestPost(w, r)
	case "PUT":
	default:
	}
}

func UsersRequestAdmin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		UsersRequestGet(w, r)
	case "POST":
		UsersRequestPostAdmin(w, r)
	case "PUT":
	case "DELETE":
		UsersRequestDeleteAdmin(w, r)
	default:
	}
}

func DataGameRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		DataGameRequestGet(w, r)
	case "POST":
		//DataGameRequestPost(w, r)
	case "PUT":
		DataGameRequestUpdate(w, r)
	default:
	}
}

func DataGameRequestAdmin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		DataGameRequestGet(w, r)
	case "POST":
		DataGameRequestPostAdmin(w, r)
	case "PUT":
		DataGameRequestUpdateAdmin(w, r)
	case "DELETE":
		DataGameRequestDeleteAdmin(w, r)
	default:
	}
}

func RegisterRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		RegisterRequestGet(w, r)
	case "POST":
		RegisterRequestPost(w, r)
	case "PUT":
		RegisterRequestUpdate(w, r)
	default:
	}
}

func RegisterRequestAdmin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		RegisterRequestGet(w, r)
	case "POST":
		RegisterRequestPostAdmin(w, r)
	case "PUT":
		RegisterRequestUpdateAdmin(w, r)
	case "DELETE":
		RegisterRequestDeleteAdmin(w, r)
	default:
	}
}

func LoginRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
		LoginRequestPost(w, r)
	case "PUT":
	default:
	}
}

func LoginRequestAdmin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
	case "PUT":
	case "DELETE":
		LoginRequestDeleteAdmin(w, r)
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

	mux := http.NewServeMux()

	mux.HandleFunc("/odisea/check", CheckServicesRequest)
	mux.HandleFunc("/odisea/register", RegisterRequest)
	mux.HandleFunc("/odisea/admin/register", RegisterRequestAdmin)
	mux.HandleFunc("/odisea/data", DataGameRequest)
	mux.HandleFunc("/odisea/admin/data", DataGameRequestAdmin)
	mux.HandleFunc("/odisea/login", LoginRequest)
	mux.HandleFunc("/odisea/admin/login", LoginRequestAdmin)
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

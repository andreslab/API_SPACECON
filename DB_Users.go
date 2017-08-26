package main

import (
	"database/sql"
	"fmt"
	"log"
)

/*
import (
	"database/sql"
	"fmt"
	"log"
)



func CreateTableUsers() {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: CONEXIÓN EXITOSA")
	}

	resp, err := db.Query("CREATE TABLE users ()")
	if err != nil {
		fmt.Println("ERROR: CREACIÓN DE TABLA")
	} else {
		fmt.Println("SUCCESS: CREACIÓN DE TABLA")
	}
	defer resp.Close()
	defer db.Close()
}

func InsertTableUsers(datatable *UsersController) {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: CONEXIÓN EXITOSA")
	}

	resp, err := db.Prepare("INSERT INTO users () VALUES ();")
	resp.Query(
		datatable.ID,
		datatable.USERNAME,
		datatable.PHONE,
		datatable.CREATED,
		datatable.DIAMONDS,
		datatable.STATE)

	if err != nil {
		fmt.Println("ERROR: INGRESO DE DATOS")
	} else {
		fmt.Println("SUCCESS")
	}
	defer resp.Close()
	defer db.Close()

}
*/

func SelectTableUserForLogin(user string, pass string) bool {
	var passw string
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Printf("ERROR: CONEXIÓN A BASE DE DATOS")
	} else {
		log.Printf("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	defer db.Close()
	err = db.QueryRow("SELECT password FROM user WHERE username = ?", user).Scan(&passw)
	if err != nil {
		fmt.Printf("ERROR: CONSULTA DE DATOS")
	} else {
		fmt.Printf("SUCCESS: CONSULTA DE DATOS")
	}
	if passw == pass {
		log.Printf("SUCCESS LOGIN")
		return true
	} else {
		log.Printf("SUCCESS LOGIN")
		return false
	}
}

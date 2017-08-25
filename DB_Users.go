package main

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

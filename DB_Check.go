package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	//_ "github.com/go-sql-driver/mysql"
)

func CreateTableCheck() {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: Conexión exitosa")
	}

	resp, _err := db.Query("create table runserver (id int NOT NULL PRIMARY KEY, state varchar(255), super varchar(255), message varchar(255), points varchar(255));")
	if _err != nil {
		fmt.Println("error en la creación de la tabla")
		log.Fatal(_err)
	} else {
		fmt.Println("creación de tabla exitosamente")
	}

	resp.Close()
}

func SelectLastDataTableCheck() {
	data := NewCheckServicesControllerEmpty()
	var index = 0
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Printf("ERROR: CONEXCIÓN A BASE DE DATOS")
	} else {
		log.Printf("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	resp, err := db.Query("SELECT * FROM runserver")
	if err != nil {
		log.Printf("ERROR: CONSULTA DE DATOS")
	} else {
		log.Printf("SUCCESS: CONSULTA DE DATOS")
	}

	defer db.Close()

	for resp.Next() {
		err := resp.Scan(&index, &data.STATE, &data.SUPER, &data.MESSAGE, &data.POINTS)
		if err != nil {
			log.Printf("ERROR: DATOS EXTRAIDOS")
		} else {
			log.Printf("SUCCESS: DATOS EXTRAIDOS")
		}
	}

	responseDataCheckService["0"] = *data

}

func SelectLastIDTableCheck() {
	var lastindex = 0
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Printf("ERROR: CONEXCIÓN A BASE DE DATOS")
	} else {
		log.Printf("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	resp, err := db.Query("SELECT MAX(id) AS maxid FROM runserver")
	if err != nil {
		log.Printf("ERROR: CONSULTA DE DATOS")
	} else {
		log.Printf("SUCCESS: CONSULTA DE DATOS")
	}

	defer resp.Close()
	defer db.Close()

	for resp.Next() {
		err := resp.Scan(&lastindex)
		if err != nil {
			log.Printf("ERROR: DATOS EXTRAIDOS")
		} else {
			log.Printf("SUCCESS: DATOS EXTRAIDOS")
			idCheckService = lastindex
			log.Printf("index: %d", idCheckService)
		}
	}

}

func InsertTableCheck(datatable *CheckServicesController) {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	fmt.Println(datatable)
	//statement : declaración
	stmt, err := db.Prepare("INSERT INTO runserver (id, state, super, message, points) VALUES (?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	index, err := strconv.Atoi(datatable.ID)

	if err != nil {
		log.Printf("ERROR: CONVERTIR INDEX")
		index = 0
	}

	_, err = stmt.Exec(
		index,
		datatable.STATE,
		datatable.SUPER,
		datatable.MESSAGE,
		datatable.POINTS)

	if err != nil {
		fmt.Println("ERROR: INGRESO DE DATOS")
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: INGRESO DE DATOS")
	}
	//defer resp.Close()
	defer db.Close()
}

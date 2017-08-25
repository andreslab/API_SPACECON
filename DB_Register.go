package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

/*func CreateTableRegister() {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: Conexión exitosa")
	}

	resp, err := db.Query("CREATE TABLE register")
	if err != nil {
		fmt.Println("ERROR: CREACIÓN DE TABLA")
	} else {
		fmt.Println("SUCCESS: CREACIÓN DE TABLA")
	}
	defer resp.Close()
	defer db.Close()
}*/

func CreateTableRegister() {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: Conexión exitosa")
	}

	resp, _err := db.Query("create table user (id int NOT NULL PRIMARY KEY, phone varchar(255), username varchar(255), created varchar(255), password varchar(255));")
	if _err != nil {
		fmt.Println("error en la creación de la tabla")
		log.Fatal(_err)
	} else {
		fmt.Println("creación de tabla exitosamente")
	}

	resp.Close()
}

func SelectTableRegister() {
	data := NewRegisterControllerEmpty()
	lastindex := 0
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Printf("ERROR: CONEXIÓN A BASE DE DATOS")
	} else {
		log.Printf("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	resp, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Printf("ERROR: CONSULTA DE DATOS")
	} else {
		fmt.Printf("SUCCESS: CONSULTA DE DATOS")
	}
	defer resp.Close()
	defer db.Close()

	for resp.Next() {
		err := resp.Scan(&lastindex, &data.PHONE, &data.USERNAME, &data.PASSWORD, &data.CREATED)
		if err != nil {
			fmt.Printf("ERROR: DATOS EXTRAIDOS")
		} else {
			fmt.Printf("SUCCESS: DATOS EXTRAIDOS")
		}
		data.ID = strconv.Itoa(lastindex)

		responseDataRegister[data.ID] = *data

	}

	err = resp.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func SelectLastIdTableRegister() int {

	lastindex := 0
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Printf("ERROR: CONEXIÓN A BASE DE DATOS")
	} else {
		log.Printf("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	resp, err := db.Query("SELECT id FROM user")
	if err != nil {
		fmt.Printf("ERROR: CONSULTA DE DATOS")
	} else {
		fmt.Printf("SUCCESS: CONSULTA DE DATOS")
	}
	defer resp.Close()
	defer db.Close()

	for resp.Next() {
		err := resp.Scan(&lastindex)
		if err != nil {
			fmt.Printf("ERROR: DATOS EXTRAIDOS")
		} else {
			fmt.Printf("SUCCESS: DATOS EXTRAIDOS")
		}
		idRegister = lastindex

	}

	err = resp.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lastindex
}

func InsertTableRegister(datatable *RegisterController) {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: CONEXIÓN EXITOSA")
	}

	resp, err := db.Prepare("INSERT INTO user (id, phone, username, password, created) VALUES (?,?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}

	index, err := strconv.Atoi(datatable.ID)

	if err != nil {
		log.Printf("ERROR: CONVERTIR INDEX")
		index = 0
	}

	_, err = resp.Exec(
		index,
		datatable.PHONE,
		datatable.USERNAME,
		datatable.PASSWORD,
		datatable.CREATED)

	if err != nil {
		fmt.Println("ERROR: INGRESO DE DATOS")
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: INGRESO DE DATOS")
	}
	//defer resp.Close()
	defer db.Close()
}

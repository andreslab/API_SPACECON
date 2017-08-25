package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func CreateTableDataGame() {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: Conexión exitosa")
	}

	resp, _err := db.Query("create table data (id int NOT NULL PRIMARY KEY, name varchar(255), value varchar(255), latitude varchar(255), longitude varchar(255), state varchar(255), created varchar(255), hunter varchar(255));")
	if _err != nil {
		fmt.Println("error en la creación de la tabla")
		log.Fatal(_err)
	} else {
		fmt.Println("creación de tabla exitosamente")
	}

	resp.Close()
}

func SelectTableData() *DataGameController {

	lastindex := 0

	data := NewDataGameControllerEmpty()
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
		fmt.Println("ERROR: CONEXIÓN A BASE DE DATOS")
	} else {
		fmt.Println("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	resp, err := db.Query("SELECT * FROM data")
	if err != nil {
		fmt.Println("ERROR: CONSULTA DE DATOS")
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: CONSULTA DE DATOS")
	}
	defer resp.Close()
	defer db.Close()

	for resp.Next() {
		err := resp.Scan(&lastindex, &data.NAME, &data.VALUE, &data.LATITUDE, &data.LONGITUDE, &data.STATE, &data.CREATED, &data.HUNTER)
		if err != nil {
			fmt.Println("ERROR: DATOS EXTRAIDOS")
			log.Fatal(err)
		} else {
			fmt.Println("SUCCESS: DATOS EXTRAIDOS")
			//log.Println(data.NAME)
		}

		data.ID = strconv.Itoa(lastindex)

		responseDataDataGame[data.ID] = *data
	}
	err = resp.Err()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

//last id

func SelectLastIdTableData() int {

	lastindex := 0

	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
		fmt.Println("ERROR: CONEXIÓN A BASE DE DATOS")
	} else {
		fmt.Println("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	resp, err := db.Query("SELECT id FROM data")
	if err != nil {
		fmt.Println("ERROR: CONSULTA DE DATOS")
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: CONSULTA DE DATOS")
	}
	defer resp.Close()
	defer db.Close()

	for resp.Next() {
		err := resp.Scan(&lastindex)
		if err != nil {
			fmt.Println("ERROR: DATOS EXTRAIDOS")
			log.Fatal(err)
		} else {
			fmt.Println("SUCCESS: DATOS EXTRAIDOS")
			//log.Println(data.NAME)
		}

		idDataGame = lastindex

	}
	err = resp.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lastindex
}

/*func CreateTableData() {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
		fmt.Println("ERROR: CONEXIÓN A BASE DE DATOS")
	} else {
		fmt.Println("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	defer db.Close()

	resp, err := db.Query("CREATE TABLE data (id int NOT NULL PRIMARY KEY, name VARCHAR(255), value VARCHAR(255),latitude VARCHAR(255),longitude VARCHAR(255),state VARCHAR(255),created VARCHAR(255),hunter VARCHAR(255));")
	if err != nil {
		fmt.Println("ERROR: CREACIÓN DE LA TABLA")
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: CREACIÓN DE LA TABLA")
	}
	defer resp.Close()

}*/

func InsertTableData(datatable *DataGameController) {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	fmt.Println(datatable)
	//statement : declaración
	stmt, err := db.Prepare("INSERT INTO data (id, name, value, latitude, state, created, hunter, longitude) VALUES (?,?,?,?,?,?,?,?);")
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
		datatable.NAME,
		datatable.VALUE,
		datatable.LATITUDE,
		datatable.LONGITUDE,
		datatable.STATE,
		datatable.CREATED,
		datatable.HUNTER)

	if err != nil {
		fmt.Println("ERROR: INGRESO DE DATOS")
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: INGRESO DE DATOS")
	}
	//defer resp.Close()
	defer db.Close()
}

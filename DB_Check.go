package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SelectTableCheck() *CheckServicesController {
	data := NewCheckServicesControllerEmpty()
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Printf("ERROR: CONEXCIÓN A BASE DE DATOS")
	} else {
		log.Printf("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	resp, err := db.Query("SELECT * FROM check")
	if err != nil {
		log.Printf("ERROR: CONSULTA DE DATOS")
	} else {
		log.Printf("SUCCESS: CONSULTA DE DATOS")
	}

	defer resp.Close()
	defer db.Close()

	for resp.Next() {
		err := resp.Scan(data.ID, data.STATE, data.SUPER, data.MESSAGE, data.POINTS)
		if err != nil {
			log.Printf("ERROR: DATOS EXTRAIDOS")
		} else {
			log.Printf("SUCCESS: DATOS EXTRAIDOS")
		}
	}

	err = resp.Err()
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func UpdateTableCheck(datatable *CheckServicesController) {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	log.Printf(datatable.STATE)

	/*
			UPDATE table_name
		SET column1 = value1, column2 = value2, ...
		WHERE condition;
	*/

	/*
			UPDATE Customers
		SET ContactName = 'Alfred Schmidt', City= 'Frankfurt'
		WHERE CustomerID = 1;
	*/

	stmt, err := db.Prepare("UPDATE check SET state = '?',super='?',message='?', points='?' WHERE id=0")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(
		datatable.STATE,
		datatable.SUPER,
		datatable.MESSAGE,
		datatable.POINTS)

	if err != nil {
		log.Panicf("ERROR: ACTUALIZACIÓN DE DATOS")
	} else {
		log.Printf("SUCCESS: INGRESO DE DATOS")
	}
}

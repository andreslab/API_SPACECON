package servicedata

import (
	"database/sql"
	"fmt"
	"log"

	controller "../../controller"
	_ "github.com/go-sql-driver/mysql"
)

func SelectTableData() *controller.DataGameController {

	data := controller.NewDataGameControllerEmpty()
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
		err := resp.Scan(&data.ID, &data.NAME, &data.VALUE, &data.LATITUDE, &data.LONGITUDE, &data.STATE, &data.CREATED, &data.HUNTER)
		if err != nil {
			fmt.Println("ERROR: DATOS EXTRAIDOS")
			log.Fatal(err)
		} else {
			fmt.Println("SUCCESS: DATOS EXTRAIDOS")
			//log.Println(data.NAME)
		}

	}
	err = resp.Err()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func CreateTableData() {
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

}

func InsertTableData(datatable *controller.DataGameController) {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: CONEXIÓN A BASE DE DATOS")
	}

	fmt.Println(datatable.NAME)
	//statement : declaración
	stmt, err := db.Prepare("INSERT INTO data (id, name, value, latitude, longitude, state, created, hunter)VALUES (?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(
		datatable.ID,
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

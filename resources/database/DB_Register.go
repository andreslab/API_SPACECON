package servicedata

import (
	"database/sql"
	"fmt"
	"log"

	controller "../../controller"
	_ "github.com/go-sql-driver/mysql"
)

func CreateTableRegister() {
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
}

func InsertTableRegister(datatable *controller.RegisterController) {
	db, err := sql.Open(typeDataBase, linkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("SUCCESS: CONEXIÓN EXITOSA")
	}

	resp, err := db.Prepare("INSERT INTO register (id, phone, username, password, created) VALUES (?,?,?,?,?)")
	resp.Query(
		datatable.ID,
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
	defer resp.Close()
	defer db.Close()
}

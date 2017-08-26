package main

import "os"
import "log"

import "net/url"

//production

//var linkDataBase = user + ":" + pass + "@tcp(" + url + ":" + portdb + ")/" + namedb
var linkDataBase string

//var linkDataBase = "admin:12345678@tcp(spacecongame.cffsdiafpjwv.us-east-2.rds.amazonaws.com:3306)/spacecondb"
var typeDataBase = "mysql"
var tableDataGame = "data"
var tableLogin = "login"
var tableRegister = "register"
var tablaUsers = "user"

func config() {

	//develop
	/*user = "bfc6a71ffb843c"
	pass = "6406db85"
	url = "us-cdbr-iron-east-05.cleardb.net"
	portdb = "3306"
	namedb = "heroku_38cbea6230473ea"*/
	u := os.Getenv("CLEARDB_DATABASE_URL")

	uri, err_parse := url.Parse(u)
	if err_parse != nil {
		log.Printf("ERROR PARSING")
	}

	//user := uri.User.Username()
	user := "bfc6a71ffb843c"
	pass := "6406db85"
	portdb := uri.Port()
	host := uri.Hostname()
	namedb := uri.Path
	log.Printf(user)
	log.Printf(pass)
	log.Printf(portdb)
	log.Printf(host)
	log.Printf(namedb)

	linkDataBase = user + ":" + pass + "@tcp(" + host + ":" + portdb + ")/" + namedb

}

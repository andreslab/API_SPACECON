package main

//production

//var linkDataBase = user + ":" + pass + "@tcp(" + url + ":" + portdb + ")/" + namedb
var linkDataBase string

//var linkDataBase = "admin:12345678@tcp(spacecongame.cffsdiafpjwv.us-east-2.rds.amazonaws.com:3306)/spacecondb"
var typeDataBase = "mysql"

func config() {

	//develop
	/*user = "bfc6a71ffb843c"
	pass = "6406db85"
	url = ""
	portdb = "3306"
	namedb = "heroku_38cbea6230473ea"*/

	/*u := os.Getenv("CLEARDB_DATABASE_URL")
	uri, err_parse := url.Parse(u)
	if err_parse != nil {
		log.Printf("ERROR PARSING")
	}*/

	//user := uri.User.Username()
	user := "bfc6a71ffb843c"
	pass := "6406db85"
	//portdb := uri.Port()
	portdb := "3306"
	host := "us-cdbr-iron-east-05.cleardb.net"
	namedb := "heroku_38cbea6230473ea"

	linkDataBase = user + ":" + pass + "@tcp(" + host + ":" + portdb + ")/" + namedb

}

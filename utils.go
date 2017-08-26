package main

const user = "bfc6a71ffb843c"
const pass = "6406db85"
const url = "us-cdbr-iron-east-05.cleardb.net"
const port = "3306"
const namedb = "heroku_38cbea6230473ea"

var linkDataBase = user + ":" + pass + "@tcp(" + url + ":" + port + ")/" + namedb

//var linkDataBase = "admin:12345678@tcp(spacecongame.cffsdiafpjwv.us-east-2.rds.amazonaws.com:3306)/spacecondb"
var typeDataBase = "mysql"
var tableDataGame = "data"
var tableLogin = "login"
var tableRegister = "register"
var tablaUsers = "user"

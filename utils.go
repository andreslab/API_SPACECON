package main

var url_heroku = "mysql://bfc6a71ffb843c:6406db85@us-cdbr-iron-east-05.cleardb.net/heroku_38cbea6230473ea?reconnect=true"

const user = "bfc6a71ffb843c"
const pass = "6406db85"

const url = "us-cdbr-iron-east-05.cleardb.net"

//const url = "127.0.0.1"
const portdb = "3306"
const namedb = "heroku_38cbea6230473ea"

//develop
//var linkDataBase = user + ":" + pass + "@tcp(" + url + ":" + portdb + ")/" + namedb

//deploy
var linkDataBase = url_heroku

//var linkDataBase = "admin:12345678@tcp(spacecongame.cffsdiafpjwv.us-east-2.rds.amazonaws.com:3306)/spacecondb"
var typeDataBase = "mysql"
var tableDataGame = "data"
var tableLogin = "login"
var tableRegister = "register"
var tablaUsers = "user"

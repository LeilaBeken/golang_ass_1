package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

var id string
var name string
var email string
var balance float64
var price float64

const (
	host     = ""
	port     = 5042
	user     = "postgres"
	password = ""
	dbname   = ""
)

var Db *sql.DB

func Init() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	Db, err = sql.Open("postgres", psqlconn)
	CheckError(err)
}


func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
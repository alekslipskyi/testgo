package db

import (
	"database/sql"
	"fmt"
	"os"
)

type Instance struct {
	name string
}

var dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"))

var connect, _ = sql.Open("postgres", dbinfo)

func (instance *Instance) findById(id string) {
	connect.QueryRow("SELECT * FROM $1", instance.name, id)
}

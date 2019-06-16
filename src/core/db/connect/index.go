package connect

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

var Dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"))

func Init() {
	connection, err := sql.Open("postgres", Dbinfo)
	if err != nil {
		panic(err)
	}

	DB = connection
}

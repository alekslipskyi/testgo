package main

import (
	"./src/controllers/user"
	"controllers/auth"
	"controllers/list"
	"core/db/connect"
	"lib/Router"
	"log"
	"net/http"
)

func main() {
	router := Router.New{}
	connect.Init()

	user.Routes()
	auth.Routes()
	list.Routes()

	http.HandleFunc("/", router.EntryPoint)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

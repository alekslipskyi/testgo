package main

import (
	"./src/controllers/user"
	"controllers/auth"
	"lib/Router"
	"log"
	"net/http"
)

func main() {
	router := Router.New{}

	user.Routes()
	auth.Routes()

	http.HandleFunc("/", router.EntryPoint)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

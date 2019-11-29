package main

import (
	"./src"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", router.Handler()))
}

package main

import (
	"log"
	"net/http"

	"github.com/cclulu/code_assign/jokes"
	"github.com/cclulu/code_assign/middleware"
)

func main() {

	address := ":5000"
	log.Println("Starting server on address", address)

	http.HandleFunc("/", middleware.Logging(jokes.Build))

	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}

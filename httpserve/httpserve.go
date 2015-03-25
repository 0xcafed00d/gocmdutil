package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := "8080"

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir("."))))
}

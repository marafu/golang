package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

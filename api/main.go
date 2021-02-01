package main

import (
	"net/http"
)

func main() {

	http.Handle("/")

	http.ListenAndServe(":8080", nil)
}

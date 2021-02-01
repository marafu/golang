package create

import (
	"fmt"
	"net/http"
)

func create(response http.ResponseWriter, request *http.Request) {
	return fmt.Fprintf(response, "Hello World")
}

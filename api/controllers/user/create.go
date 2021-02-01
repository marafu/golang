package create

import (
	"fmt"
	"net/http"
)

func create(response http.ResponseWriter, request *http.Request) string {
	return fmt.Fprintf(response, "Hello World")
}

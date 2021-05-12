package api

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, response *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello World!!")
}

package middlewares

import (
	"log"
	"net/http"
)

// Log of requests
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			log.Printf("%s. IP: %s. URL : %s\n",
				request.Method, request.RemoteAddr, request.URL)
			handler.ServeHTTP(response, request)
		})
}

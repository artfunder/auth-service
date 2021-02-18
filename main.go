package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CLIENT_SITE"))
	fmt.Fprintf(w, "Hello, world")
}

func main() {
	http.Handle("/", http.HandlerFunc(handler))

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

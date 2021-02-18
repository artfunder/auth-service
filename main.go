package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CLIENT_SITE"))
}

func main() {
	http.Handle("/", http.HandlerFunc(handler))

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil)
}

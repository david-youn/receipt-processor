package main

import (
	"net/http"

	"receipt-processor/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}

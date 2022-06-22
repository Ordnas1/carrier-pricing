package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ordnas1/carrier-pricing/internal/services/quotes"
)

func helloWorldHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloWorldHandle)
	mux.HandleFunc("/quotes", quotes.QuoteHandler)

	fmt.Println("Listening")
	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		log.Fatal(err)
	}
}

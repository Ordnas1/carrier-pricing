package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloWorldHandle)
}

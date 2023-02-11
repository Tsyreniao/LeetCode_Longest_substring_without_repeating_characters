package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)

	fmt.Println("Listening on port :3000...")

	http.ListenAndServe(":3000", nil)
}
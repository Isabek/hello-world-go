package main

import (
	"net/http"
	"fmt"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", Hello)

	http.ListenAndServe(":8080", nil)
}

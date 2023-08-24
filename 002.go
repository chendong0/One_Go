package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "陈东 Hello, World!")
	})
	http.ListenAndServe("localhost:8080", nil)
}

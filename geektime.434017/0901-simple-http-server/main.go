package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	fmt.Printf("Server started on %s .\n", "http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

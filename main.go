package main

import (
	"fmt"
	"net/http"
)

func HelloGolang(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Golang")
}

func main() {
	http.HandleFunc("/", HelloGolang)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

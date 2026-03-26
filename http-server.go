package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello Thái")
}

func main() {

	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}

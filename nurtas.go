package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello world!")
	})

	log.Fatal(http.ListenAndServe(":3333", nil))
}

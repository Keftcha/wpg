package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", gallery)

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

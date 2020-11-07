package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", gallery)

	http.Handle(
		"/pics/",
		http.StripPrefix(
			"/pics/",
			http.FileServer(http.Dir("/pics/")),
		),
	)

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

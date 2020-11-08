package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/keftcha/wpg/handlers"
)

func main() {
	http.HandleFunc("/", handelers.Gallery)
	http.HandleFunc("/slideshow/", handelers.Slideshow)

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

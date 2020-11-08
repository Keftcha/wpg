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

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static/")),
		),
	)

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

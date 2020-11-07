package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func gallery(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("pages/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	tpl.Execute(w, struct{}{})
}

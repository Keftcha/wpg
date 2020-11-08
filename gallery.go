package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func gallery(w http.ResponseWriter, r *http.Request) {
	// Get the directory where the user want to be
	path := r.URL.Path
	// Read files in the /pics directory
	files, err := ioutil.ReadDir("/pics" + path)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	// Separate directories and pictures
	dirs, pics := make([]string, 0), make([]string, 0)
	for _, file := range files {
		// Check if it's a directory
		if file.IsDir() {
			dirs = append(dirs, path+file.Name())
		} else {
			pics = append(pics, path+file.Name())
		}
	}

	// Make the map with infos to parse the template
	info := make(map[string][]string)
	info["dirs"] = dirs
	info["pics"] = pics

	// Format and send the template page
	tpl, err := template.ParseFiles("pages/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	tpl.Execute(w, info)
}

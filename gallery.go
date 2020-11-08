package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type webFile struct {
	Name string
	Path string
}

func gallery(w http.ResponseWriter, r *http.Request) {
	// Get the directory where the user want to be
	path := r.URL.Path

	// Read files in the /pics directory
	files, err := ioutil.ReadDir("/pics" + path)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	// Separate directories and pictures
	dirs, pics := make([]webFile, 0), make([]webFile, 0)
	if path != "/" {
		// Add the root directory
		dirs = append(dirs, webFile{Name: "/", Path: "/"})
		// Add the parent directory
		dirs = append(
			dirs,
			webFile{
				Name: "../",
				Path: parentDir(path) + "/",
			},
		)
	}
	// Loop over directory content
	for _, file := range files {
		// Make the webFile
		wbfle := webFile{
			Name: file.Name(),
			Path: path + file.Name(),
		}
		// Check if it's a directory
		if file.IsDir() {
			wbfle.Path = wbfle.Path + "/"
			dirs = append(dirs, wbfle)
		} else {
			wbfle.Path = "/pics" + wbfle.Path
			pics = append(pics, wbfle)
		}
	}

	// Make the map with infos to parse the template
	info := make(map[string][]webFile)
	info["dirs"] = dirs
	info["pics"] = pics

	// Format and send the template page
	tpl, err := template.ParseFiles("pages/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	tpl.Execute(w, info)
}

func parentDir(path string) string {
	// Remove trailing `/` caracter
	path = path[:len(path)-1]
	// Find the index of the last `/`
	idx := strings.LastIndex(path, "/")
	// Remove all thing after the last `/`
	return path[:idx]
}

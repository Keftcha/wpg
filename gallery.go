package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func gallery(w http.ResponseWriter, r *http.Request) {
	// Get the directory where the user want to be
	path := r.URL.Path

	// Read files in the /pics directory
	files, err := ioutil.ReadDir("/pics" + path)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	// Make the map with infos to parse the template and
	// Separate directories and pictures
	info := make(map[string]interface{}) // ← This is realy bad
	info["crntPath"] = path
	info["dirs"], info["pics"] = distinctDirsAndPics(files, path)

	// Format and send the template page
	tpl, err := template.ParseFiles("pages/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	tpl.Execute(w, info)
}

func newWebFileDir(file os.FileInfo, path string) webFile {
	return webFile{
		Name: file.Name(),
		Path: path + file.Name() + "/",
	}
}

func newWebFileImage(file os.FileInfo, path string) webFile {
	return webFile{
		Name: file.Name(),
		Path: "/pics" + path + file.Name(),
	}
}

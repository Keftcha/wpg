package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
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
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	// Make the map with infos to parse the template and
	// Separate directories and pictures
	info := make(map[string][]webFile)
	info["dirs"], info["pics"] = distinctDirsAndPics(files, path)

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

func distinctDirsAndPics(files []os.FileInfo, path string) ([]webFile, []webFile) {
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
		// Check if it's a directory
		if file.IsDir() {
			dirs = append(dirs, newWebFileDir(file, path))
		} else {
			pics = append(pics, newWebFileImage(file, path))
		}
	}

	return dirs, pics
}

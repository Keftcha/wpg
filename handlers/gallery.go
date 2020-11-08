package handelers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/keftcha/wpg/helpers"
)

func Gallery(w http.ResponseWriter, r *http.Request) {
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
	info := make(map[string]interface{})
	info["crntPath"] = path
	info["dirs"], info["pics"] = helpers.DistinctDirsAndPics(files, path)

	// Format and send the template page
	tpl, err := template.ParseFiles("pages/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	tpl.Execute(w, info)
}

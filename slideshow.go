package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func slideshow(w http.ResponseWriter, r *http.Request) {
	// Query parameters:
	// - type → slideshow type
	// - img → current img path displayed
	// - delay → delay before changing img in sec
	// - dir → directory of the slideshow

	// Get slide show type
	// - crt → Current directory
	// - sub → Current and sub-directories
	// - all → All pictures found
	slideshowType := r.URL.Query().Get("type")
	if !isValidSlideshowType(slideshowType) {
		slideshowType = "crt"
	}

	// Get the delay time
	delay := validateDelay(r.URL.Query().Get("delay"))

	// Get the directory
	dir := r.URL.Query().Get("dir")
	if dir == "" {
		dir = "/"
	}
	// Check if it exist
	if _, err := os.Stat("/pics" + dir); err != nil {
		http.Error(
			w,
			fmt.Sprintf("Directory %s doesn't exist", dir),
			404,
		)
		return
	}

	// Define images names list
	filesName := make([]webFile, 0)
	switch slideshowType {
	case "crt":
		// Read files in the asked directory (prefixed by `/pics`)
		files, err := ioutil.ReadDir("/pics" + dir)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusInternalServerError)
			return
		}
		// Get pictures file
		_, filesName = distinctDirsAndPics(files, dir)
	case "sub":
		reccursivlyFindPics(dir, &filesName)
	case "all":
		reccursivlyFindPics("/", &filesName)
	}

	if len(filesName) <= 0 {
		http.Error(w, "No image found.", 404)
		return
	}

	// Get the image to display
	imgName := r.URL.Query().Get("img")
	imgIdx := index(filesName, imgName)
	var img webFile
	if imgName == "" || imgIdx == -1 {
		img = (filesName)[0]
		imgIdx = 0
	} else {
		img = (filesName)[imgIdx]
	}

	// Get the next image file path
	nextImgPath := (filesName)[(imgIdx+1)%len(filesName)].Path

	// Regroup info
	info := map[string]string{
		"type":        slideshowType,
		"imgPath":     img.Path,
		"imgName":     img.Name,
		"nextImgPath": nextImgPath,
		"delay":       fmt.Sprint(delay),
		"dir":         dir,
	}
	// Parse template
	tpl, err := template.ParseFiles("pages/slideshow.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	tpl.Execute(w, info)

}

func isValidSlideshowType(sst string) bool {
	return sst == "crt" || sst == "sub" || sst == "all"
}

func validateDelay(delay string) int {
	nb, err := strconv.ParseInt(delay, 10, 64)
	if err != nil {
		return 5
	}
	return int(nb)
}

func index(lst []webFile, str string) int {
	for idx := 0; idx < len(lst); idx++ {
		if lst[idx].Path == str {
			return idx
		}
	}
	return -1
}

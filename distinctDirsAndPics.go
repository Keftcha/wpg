package main

import (
	"os"
)

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

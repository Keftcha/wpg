package helpers

import (
	"os"
)

func DistinctDirsAndPics(files []os.FileInfo, path string) ([]WebFile, []WebFile) {
	dirs, pics := make([]WebFile, 0), make([]WebFile, 0)
	if path != "/" {
		// Add the root directory
		dirs = append(dirs, WebFile{Name: "/", Path: "/"})
		// Add the parent directory
		dirs = append(
			dirs,
			WebFile{
				Name: "../",
				Path: ParentDir(path) + "/",
			},
		)
	}
	// Loop over directory content
	for _, file := range files {
		// Check if it's a directory
		if file.IsDir() {
			dirs = append(dirs, NewWebFileDir(file, path))
		} else {
			pics = append(pics, NewWebFileImage(file, path))
		}
	}

	return dirs, pics
}

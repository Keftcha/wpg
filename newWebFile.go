package main

import (
	"os"
)

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

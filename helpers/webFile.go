package helpers

import (
	"os"
)

type WebFile struct {
	Name string
	Path string
}

func NewWebFileDir(file os.FileInfo, path string) WebFile {
	return WebFile{
		Name: file.Name(),
		Path: path + file.Name() + "/",
	}
}

func NewWebFileImage(file os.FileInfo, path string) WebFile {
	return WebFile{
		Name: file.Name(),
		Path: "/pics" + path + file.Name(),
	}
}

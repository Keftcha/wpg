package helpers

import (
	"strings"
)

func ParentDir(path string) string {
	// Remove trailing `/` caracter
	path = path[:len(path)-1]
	// Find the index of the last `/`
	idx := strings.LastIndex(path, "/")
	// Remove all thing after the last `/`
	if idx != -1 {
		return path[:idx]
	}
	return path
}

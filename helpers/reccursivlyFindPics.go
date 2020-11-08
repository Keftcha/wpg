package helpers

import (
	"fmt"
	"io/ioutil"
)

func ReccursivlyFindPics(baseDir string, wbfls *[]WebFile) {
	fls, err := ioutil.ReadDir("/pics" + baseDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, fle := range fls {
		if fle.IsDir() {
			// Search in the directory other images
			ReccursivlyFindPics(baseDir+fle.Name()+"/", wbfls)
		} else {
			// Add finded pictures
			newWbfls := append(*wbfls, NewWebFileImage(fle, baseDir))
			*wbfls = newWbfls
		}
	}
}

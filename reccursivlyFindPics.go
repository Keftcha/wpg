package main

import (
	"fmt"
	"io/ioutil"
)

func reccursivlyFindPics(baseDir string, wbfls *[]webFile) {
	fls, err := ioutil.ReadDir("/pics" + baseDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, fle := range fls {
		if fle.IsDir() {
			// Search in the directory other images
			reccursivlyFindPics(baseDir+fle.Name()+"/", wbfls)
		} else {
			// Add finded pictures
			newWbfls := append(*wbfls, newWebFileImage(fle, baseDir))
			*wbfls = newWbfls
		}
	}
}

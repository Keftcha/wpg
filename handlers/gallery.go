package handelers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

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
	info := struct {
		CrntPath string
		Dirs     []helpers.WebFile
		Pics     []helpers.WebFile
		NbPics   int64
		Page     int
		Pages    []map[string]int
	}{}
	info.CrntPath = path
	info.Dirs, info.Pics = helpers.DistinctDirsAndPics(files, path)

	// Define elements needs to be display with paging
	elem := r.URL.Query().Get("elem")
	page := r.URL.Query().Get("page")
	if nbPics, err := strconv.ParseInt(elem, 10, 64); err == nil && nbPics != 0 {
		pageNb, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			pageNb = 0
		}

		// Make the paging
		totPages := len(info.Pics) / int(nbPics)
		if len(info.Pics)%int(nbPics) != 0 {
			totPages++
		}
		pages := make([]map[string]int, totPages)
		for i := 0; i < totPages; i++ {
			pages[i] = map[string]int{
				"nbPics":      int(nbPics),
				"page":        i,
				"pageDisplay": i + 1,
			}
		}
		info.Pages = pages

		info.NbPics = nbPics
		info.Pics = info.Pics[pageNb*nbPics : int(math.Min(float64((pageNb+1)*nbPics), float64(len(info.Pics))))]
	}

	// Format and send the template page
	tpl, err := template.ParseFiles("pages/gallery.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = tpl.Execute(w, info)
	if err != nil {
		fmt.Println(err.Error())
	}
}

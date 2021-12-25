package helpers

type Page struct {
	NbPics        int64
	Page          int
	PageDisplay   int
	IsCurrentPage bool
}

// ListPages return the list of pages
func ListPages(totImg int, nbPics, pageNb int64) []Page {
	totPages := totImg / int(nbPics)
	if totImg%int(nbPics) != 0 {
		totPages++
	}
	pages := make([]Page, totPages)
	for i := 0; i < totPages; i++ {
		pages[i] = Page{
			NbPics:        nbPics,
			Page:          i,
			PageDisplay:   i + 1,
			IsCurrentPage: i == int(pageNb),
		}
	}

	return pages
}

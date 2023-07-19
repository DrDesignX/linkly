package main

import (
	// "net/http"
	"github.com/drdesignx/url-shortener/model"
)

var redirectMap = map[string]string{
	"/gg" : "https://google.com",
	"/twit" : "https://twitter.com",
	"/gith" : "https://github.com",
	"/digi" : "https://digikala.com",
}

// func redirectHandler(w http.ResponseWriter, r *http.Request) {
// 	// redirectURL , exists := redirectMap[r.URL.Path]
// 	// if exists {
// 	// 	http.Redirect(w, r, redirectURL, http.StatusFound)
// 	// 	return
// 	// }
// 	// http.NotFound(w,r)
// }

func main() {
	model.Setup()


}

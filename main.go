package main // import "rice-practice"

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
)

func main() {
	// http.StripPrefix should be used when using sub-path
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./not-bundled"))))
	http.Handle("/www/", http.StripPrefix("/www/", http.FileServer(rice.MustFindBox("www").HTTPBox())))

	http.ListenAndServe(":8080", nil)
}

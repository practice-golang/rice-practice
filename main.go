package main // import "rice-practice"

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
)

func main() {
	http.Handle("/", http.FileServer(rice.MustFindBox("www").HTTPBox()))
	http.ListenAndServe(":8080", nil)
}

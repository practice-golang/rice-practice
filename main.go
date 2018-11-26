package main // import "rice-practice"

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
)

func main() {
	// router := mux.NewRouter()
	// router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("www").HTTPBox()))
	// log.Fatal(http.ListenAndServe(":12345", router))

	http.Handle("/", http.FileServer(rice.MustFindBox("www").HTTPBox()))
	http.ListenAndServe(":8080", nil)
}

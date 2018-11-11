package webui

import (
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	root := rice.MustFindBox("static").HTTPBox()

	router := mux.NewRouter()
	router.PathPrefix("/static").Handler(http.FileServer(root))
	log.Fatal(http.ListenAndServe(":18080", router))
}

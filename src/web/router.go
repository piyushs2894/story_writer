package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (web *Web) InitHandlers() {
	// Word
	r := mux.NewRouter()

	r.Handle("/add", HandlerFunc(web.AddWord))
	r.Handle("/stories", HandlerFunc(web.GetStories))
	r.Handle("/stories/{id:[0-9]+}", HandlerFunc(web.GetStoryById))

	http.Handle("/", r)
}

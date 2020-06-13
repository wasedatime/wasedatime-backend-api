package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/comment/{course_id}", courseCommentHandler).Methods(http.MethodGet)
}

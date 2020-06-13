package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/comment/{course_id}", courseCommentHandler).Methods(http.MethodGet)
	log.Println("Started api service at port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

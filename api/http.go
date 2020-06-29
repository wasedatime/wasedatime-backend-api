package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).Methods(http.MethodGet)
	r.HandleFunc("/evaluations", courseEvalHandler).Methods(http.MethodGet)
	log.Println("Started api service at port 6001")
	err := http.ListenAndServe(":6001", r)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

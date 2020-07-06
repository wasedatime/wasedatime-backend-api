package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", " text/plain")
	n, err := fmt.Fprintf(w, "WasedaTime is Back!")
	if err != nil {
		log.Println(n, err)
	}
}

func courseCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Referer") != API_URL {
		w.WriteHeader(http.StatusForbidden)
		log.Println("[*Cross-Origin*]: User-Agent: " + r.Header.Get("User-Agent") + " IP: " + r.RemoteAddr)
		return
	}
	w.Header().Set("Content-Type", " application/json")
	w.Header().Set("Referrer-Policy", "origin")
	w.Header().Set("Access-Control-Allow-Origin", API_URL)
	params := r.URL.Query()
	courseId := params.Get("course_code")
	resp := findCourseCommentByID(courseId)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

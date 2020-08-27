package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", " text/plain")
	n, err := fmt.Fprintf(w, "WasedaTime is Back!")
	if err != nil {
		log.Println(n, err)
	}
}

func courseEvalTestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", " application/json")
	params := r.URL.Query()
	courseKey := params.Get("course_key")
	resp := findCourseEvalByCourseKey(courseKey, COURSE_EVAL_TEST_COLLECTION)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

func courseEvalHandler(w http.ResponseWriter, r *http.Request) {
	referer := r.Header.Get("Referer")
	refererUrl, err_ := url.Parse(referer)
	if refererUrl.Host != origin {
		w.WriteHeader(http.StatusForbidden)
		log.Println("[*Cross-Origin*]: User-Agent: " + r.Header.Get("User-Agent") + " IP: " + r.RemoteAddr)
		return
	}
	w.Header().Set("Content-Type", " application/json")
	w.Header().Set("Referrer-Policy", "origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	params := r.URL.Query()
	courseKey := params.Get("course_key")
	resp := findCourseEvalByCourseKey(courseKey, COURSE_EVAL_COLLECTION)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil || err_ != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

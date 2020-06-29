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

func courseEvalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", " application/json")
	params := r.URL.Query()
	courseCode := params.Get("course_code")
	resp := findCourseEvalByCourseCode(courseCode)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

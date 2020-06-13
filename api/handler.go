package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func courseCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", " application/json")
	courseId := mux.Vars(r)["course_id"]
	resp := findCourseCommentByID(courseId)
	if resp.CourseId == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

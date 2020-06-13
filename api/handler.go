package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func courseCommentHandler(w http.ResponseWriter, r *http.Request) {
	courseId := mux.Vars(r)["course_id"]
	resp := findCourseCommentByID(courseId)
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

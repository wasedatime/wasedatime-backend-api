package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func courseCommentHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["comments"]
	courseId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp := findCourseCommentByID(courseId)
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

package api

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CourseEval struct {
	Id            primitive.ObjectID `bson:"_id" json:"-"`
	CourseKey     string             `bson:"course_key" json:"course_key"`
	CourseCode    string             `bson:"course_code" json:"course_code"`
	Title         string             `json:"title"`
	Instructor    string             `json:"instructor"`
	Satisfaction  int8               `json:"satisfaction"`
	Difficulty    int8               `json:"difficulty"`
	Benefit       int8               `json:"benefit"`
	Comment       string             `json:"comment"`
	CommentedDate time.Time          `json:"commented_date"`
}

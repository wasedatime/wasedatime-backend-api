package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type CourseComment struct {
	Id            primitive.ObjectID `bson:"_id" json:"-"`
	CourseCode    string             `bson:"course_code" json:"course_code"`
	Title         string             `json:"title"`
	Instructor    string             `json:"instructor"`
	Satisfaction  int8               `json:"satisfaction"`
	Difficulty    int8               `json:"difficulty"`
	Benefit       int8               `json:"benefit"`
	CommentedYear int16              `bson:"commented_year" json:"commented_year"`
	Comment       string             `json:"comment"`
}

package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type CourseComment struct {
	ID            primitive.ObjectID `bson:"_id" json:"-"`
	CourseKey     string             `bson:"course_key" json:"course_key"`
	CourseCode    string             `bson:"course_code" json:"course_code"`
	Title         string             `json:"title"`
	Instructor    string             `json:"instructor"`
	Satisfaction  int8               `json:"satisfaction"`
	Difficulty    int8               `json:"difficulty"`
	Benefit       int8               `json:"benefit"`
	CommentedDate int16              `bson:"commented_date" json:"commented_date"`
	Comment       string             `json:"comment"`
}

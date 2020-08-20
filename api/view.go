package api

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CourseEval struct {
	Id            	primitive.ObjectID `bson:"_id" json:"-"`
	CourseKey     	string             `bson:"course_key" json:"course_key"`
	CourseCode    	string             `bson:"course_code" json:"course_code"`
	Title         	string             `json:"title"`
	TitleJp         string             `json:"title_jp"`
	Instructor    	string             `json:"instructor"`
	InstructorJp  	string             `json:"instructor_jp"`
	Satisfaction  	int8               `json:"satisfaction"`
	Difficulty    	int8               `json:"difficulty"`
	Benefit       	int8               `json:"benefit"`
	CommentSrcLng   int8               `json:"comment_src_lng"`
	CommentZh       string             `json:"comment_zh"`
	CommentEn       string             `json:"comment_en"`
	CommentJp       string			   `json:"comment_jp"`
	Year            int8               `json:"year"`
	CommentedDate   time.Time          `bson:"commented_date" json:"commented_date"`
}

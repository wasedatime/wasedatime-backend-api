package api

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CourseEval struct {
	Id            	primitive.ObjectID  `bson:"_id" json:"-"`
	CourseKey     	string              `bson:"course_key" json:"course_key"`
	CourseCode    	string              `bson:"course_code" json:"course_code"`
	Title         	string              `bson:"title" json:"title"`
	TitleJp         string              `bson:"title_jp" json:"title_jp"`
	Instructor    	string              `bson:"instructor" json:"instructor"`
	InstructorJp  	string              `bson:"instructor_jp" json:"instructor_jp"`
	Satisfaction  	int8                `json:"satisfaction"`
	Difficulty    	int8                `json:"difficulty"`
	Benefit       	int8                `json:"benefit"`
	CommentSrcLng   int8                `json:"comment_src_lng"`
	CommentZh       string              `bson:"comment_zh" json:"comment_zh"`
	CommentEn       string              `bson:"comment_en" json:"comment_en"`
	CommentJp       string              `bson:"comment_jp" json:"comment_jp"`
	Year            int16               `json:"year"`
	CommentedDate   time.Time           `bson:"commented_date" json:"commented_date"`
}

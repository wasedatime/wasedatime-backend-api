package api

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CourseEval struct {
	Id            primitive.ObjectID `bson:"_id" json:"-"`
	CourseKey     string             `bson:"course_key" json:"course_key"`
	CourseCode    string             `bson:"course_code" json:"course_code"`
	Title         string             `bson:"title" json:"title"`
	TitleJp       string             `bson:"title_jp" json:"title_jp"`
	Instructor    string             `bson:"instructor" json:"instructor"`
	InstructorJp  string             `bson:"instructor_jp" json:"instructor_jp"`
	Satisfaction  int8               `bson:"satisfaction" json:"satisfaction"`
	Difficulty    int8               `bson:"difficulty" json:"difficulty"`
	Benefit       int8               `bson:"benefit" json:"benefit"`
	CommentSrcLng int8               `bson:"comment_src_lng" json:"comment_src_lng"`
	CommentZhTw   string             `bson:"comment_zh_TW" json:"comment_zh_TW"`
	CommentZhCn   string             `bson:"comment_zh_CN" json:"comment_zh_CN"`
	CommentEn     string             `bson:"comment_en" json:"comment_en"`
	CommentJp     string             `bson:"comment_jp" json:"comment_jp"`
	Year          int16              `bson:"year" json:"year"`
	CommentedDate time.Time          `bson:"commented_date" json:"commented_date"`
}

type CourseEvals struct {
	CourseKey string        `json:"course_key"`
	Comments  []*CourseEval `json:"comments"`
}

type CourseEvalsRequest struct {
	CourseKeys []string `json:"course_keys"`
}

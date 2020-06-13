package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type CourseComment struct {
	Id              primitive.ObjectID `bson:"_id" json:"-"`
	CourseId        string             `bson:"course_id" json:"course_id"`
	Type            string             `json:"type"`
	CJLLowestLevel  int8               `bson:"cjl_lowest_level" json:"cjl_lowest_level"`
	CJLHighestLevel int8               `bson:"cjl_highest_level" json:"cjl_highest_level"`
	Title           string             `json:"title"`
	Instructor      string             `json:"instructor"`
	Term            string             `json:"term"`
	Occurrences     []Occurrence       `json:"occurrences"`
	Satisfaction    int8               `json:"satisfaction"`
	Difficulty      int8               `json:"difficulty"`
	Benefit         int8               `json:"benefit"`
	CommentedYear   int16              `json:"commented_year"`
	Comment         string             `json:"comment"`
}

type Occurrence struct {
	Day    string `json:"day"`
	Period int8   `json:"period"`
}

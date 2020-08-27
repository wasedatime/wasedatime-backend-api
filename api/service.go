package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const COURSE_KEY = "course_key"

func findCourseEvalByCourseKey(courseKey string, collection string) []*CourseEval {
	resp := make([]*CourseEval, 0)
	cur := &mongo.Cursor{}
	filter := bson.D{{COURSE_KEY, courseKey}}
	cur, _ = client.Database(COURSE_EVAL_DB).Collection(collection).Find(context.TODO(), filter)
	for cur.Next(context.TODO()) {
		courseEval := &CourseEval{}
		err := cur.Decode(courseEval)
		if err != nil {
			continue
		}
		resp = append(resp, courseEval)
	}
	return resp
}

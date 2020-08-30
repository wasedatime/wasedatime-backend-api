package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const COURSE_KEY = "course_key"

func findCourseEvalByCourseKey(courseKey string) []*CourseEval {
	resp := make([]*CourseEval, 0)
	cur := &mongo.Cursor{}
	filter := bson.D{{COURSE_KEY, courseKey}}
	cur, _ = client.Database(COURSE_EVAL_DB).Collection(COURSE_EVAL_COLLECTION).Find(context.TODO(), filter)
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

func findCourseEvals(courseKeys []string) []*CourseEvals {
	resp := make([]*CourseEvals, 0)
	cur := &mongo.Cursor{}
	for _, courseKey := range courseKeys {
		filter := bson.D{{COURSE_KEY, courseKey}}
		evals := make([]*CourseEval, 0)
		cur, _ = client.Database(COURSE_EVAL_DB).Collection(courseEval).Find(context.TODO(), filter)
		for cur.Next(context.TODO()) {
			courseEval := &CourseEval{}
			err := cur.Decode(courseEval)
			if err != nil {
				continue
			}
			evals = append(evals, courseEval)
		}
		resp = append(resp, &CourseEvals{CourseKey: courseKey, Comments: evals})
	}
	return resp
}

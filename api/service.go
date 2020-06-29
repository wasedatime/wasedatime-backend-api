package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const ID = "course_id"

func findCourseEvalByCourseCode(id string) []*CourseEval {
	resp := make([]*CourseEval, 0)
	cur := &mongo.Cursor{}
	filter := bson.D{{ID, id}}
	cur, _ = client.Database(COMMENT_DB).Collection(COMMENT_COLLECTION).Find(context.TODO(), filter)
	for cur.Next(context.TODO()) {
		comment := &CourseEval{}
		err := cur.Decode(comment)
		if err != nil {
			continue
		}
		resp = append(resp, comment)
	}
	return resp
}

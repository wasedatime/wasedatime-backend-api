package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

const ID = "course_id"

func findCourseCommentByID(id string) *CourseComment {
	filter := bson.D{{ID, id}}
	singleResult := client.Database(COMMENT_DB).Collection(COMMENT_COLLECTION).FindOne(context.TODO(), filter)
	resp := &CourseComment{}
	_ = singleResult.Decode(resp)
	return resp
}

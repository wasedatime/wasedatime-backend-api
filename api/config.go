package api

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

const (
	CourseEvalDb = "course_evals"
)

var (
	client     *mongo.Client
	origin     = os.Getenv("ORIGIN")
	courseEval = os.Getenv("COLLECTION")
)

// initialize connections to mongodb
func init() {
	client = ConnectMgo(os.Getenv("MONGO_URL"))
}

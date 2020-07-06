package api

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

const (
	MONGO_URL              = "MONGO_URL" // envvar
	COURSE_EVAL_DB         = "course_evals"
	COURSE_EVAL_COLLECTION = "test"
	API_URL                = "https://www.wasedatime.com/"
)

var (
	client *mongo.Client
)

// initialize connections to mongodb
func init() {
	client = ConnectMgo(os.Getenv(MONGO_URL))
}

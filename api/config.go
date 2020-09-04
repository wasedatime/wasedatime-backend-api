package api

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

const (
	MONGO_URL      = "MONGO_URL" // envvar
	ORIGIN         = "ORIGIN"    // envvar
	COURSE_EVAL_DB = "course_evals"
)

var (
	client *mongo.Client
	origin = os.Getenv(ORIGIN)
)

// initialize connections to mongodb
func init() {
	client = ConnectMgo(os.Getenv(MONGO_URL))
}

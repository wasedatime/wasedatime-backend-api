package api

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

const (
	MONGO_URL          = "DB_URL" // envvar
	COMMENT_DB         = "comments"
	COMMENT_COLLECTION = "comment_test"
	API_URL            = "https://www.wasedatime.com/"
)

var (
	client *mongo.Client
)

// initialize connections to mongodb
func init() {
	client = ConnectMgo(os.Getenv(MONGO_URL))
}

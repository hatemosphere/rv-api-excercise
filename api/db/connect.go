package db

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
)

var (
	Session *mgo.Session

	Mongo *mgo.DialInfo
)

const (
	LocalMongoDBUrl = "mongodb://localhost:27017/titanic"
)

func Connect() {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		log.Printf("MONGODB_URL variable is not set, falling back to localhost")
		uri = LocalMongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	if err != nil {
		log.Printf("Error parsing mongo URI string %v\n", err)
		panic(err.Error())
	}
	s, err := mgo.Dial(uri)
	if err != nil {
		log.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	log.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}

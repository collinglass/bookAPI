package db

import (
	"fmt"
	"github.com/collinglass/bookAPI/schema"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
)

var (
	mongoSession *mgo.Session
	database     *mgo.Database
	repo         bookRepo
)

type Msg struct {
	Id    bson.ObjectId `bson:"_id"`
	Msg   string        `bson:"Msg"`
	Count int           `bson:"count"`
}

type Environment struct {
	DB         string
	collection string
}

func Connect() (*mgo.Session, error) {
	uri := "mongodb://collinglass:bookAPI@troup.mongohq.com:10099/bookAPI"
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	sess, err := mgo.Dial(uri)

	sess.SetSafe(&mgo.Safe{})
	return sess, err
}

func UpdateMsg(sess *mgo.Session, db string, col string, doc *Msg, update interface{}) error {
	collection := sess.DB(db).C(col)
	err := collection.Update(bson.M{"_id": doc.Id}, update)
	return err
}

func InsertMongo(book *schema.Book) {
	var err error
	// Setup the database
	if mongoSession, err = mgo.Dial("mongodb://collinglass:bookAPI@troup.mongohq.com:10099/bookAPI"); err != nil {
		panic(err)
	}
	log.Println("Connected to mongodb")

	database = mongoSession.DB("bookAPI")
	repo.Collection = database.C("books")

	repo.Create(book)
}

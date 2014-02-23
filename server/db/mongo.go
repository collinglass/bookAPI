package db

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
)

type msg struct {
	Id    bson.ObjectId `bson:"_id"`
	Msg   string        `bson:"msg"`
	Count int           `bson:"count"`
}

type environment struct {
	DB         string
	collection string
}

func connect() (*mgo.Session, error) {
	uri := "mongodb://collinglass:bookAPI@troup.mongohq.com:10099/bookAPI"
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	sess, err := mgo.Dial(uri)

	sess.SetSafe(&mgo.Safe{})
	return sess, err
}

func insertMsg(sess *mgo.Session, db string, col string, doc *msg) error {
	collection := sess.DB(db).C(col)
	err := collection.Insert(doc)
	return err
}

func updateMsg(sess *mgo.Session, db string, col string, doc *msg, update interface{}) error {
	collection := sess.DB(db).C(col)
	err := collection.Update(bson.M{"_id": doc.Id}, update)
	return err
}

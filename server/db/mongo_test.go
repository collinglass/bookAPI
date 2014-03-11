package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"labix.org/v2/mgo/bson"
	"os"
	"testing"
)

func TestMongoInsert(t *testing.T) {
	sess, err := Connect()

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer sess.Close()

	doc := Msg{Id: bson.NewObjectId(), Msg: "Goodby from go"}
	err = InsertMsg(sess, "bookAPI", "test", &doc)
	if err != nil {
		fmt.Printf("Can't insert document: %v\n", err)
		os.Exit(1)
	}

	var mongoDoc Msg
	err = sess.DB("bookAPI").C("test").Find(bson.M{}).One(&mongoDoc)
	if err != nil {
		fmt.Printf("Got an error finding a doc %v\n")
		os.Exit(1)
	}

	assert.Equal(t, doc, mongoDoc)
}

func TestMongoUpdate(t *testing.T) {
	sess, err := Connect()

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer sess.Close()

	var doc Msg
	err = sess.DB("bookAPI").C("test").Find(bson.M{}).One(&doc)
	if err != nil {
		fmt.Printf("got an error finding a doc %v\n")
		os.Exit(1)
	}

	update := bson.M{"$inc": bson.M{"count": 1}}
	err = UpdateMsg(sess, "bookAPI", "test", &doc, update)
	if err != nil {
		fmt.Printf("Can't update document %v\n", err)
		os.Exit(1)
	}

	var mongoDoc Msg
	err = sess.DB("bookAPI").C("books").Find(bson.M{}).One(&mongoDoc)
	if err != nil {
		fmt.Printf("got an error finding a doc %v\n")
		os.Exit(1)
	}

	assert.Nil(t, err)
	assert.NotEqual(t, doc, mongoDoc)
}

func TestMongoDrop(t *testing.T) {
	sess, err := Connect()

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer sess.Close()

	err = sess.DB("bookAPI").C("test").DropCollection()
	if err != nil {
		fmt.Printf("Got an error trying to drop %v\n")
		os.Exit(1)
	}

	assert.Nil(t, err)
}

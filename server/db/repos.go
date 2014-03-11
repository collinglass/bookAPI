package db

import (
	"github.com/collinglass/bookAPI/schema"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type (
	bookRepo struct {
		Collection *mgo.Collection
	}
)

func (r bookRepo) All() (books schema.Books, err error) {
	err = r.Collection.Find(bson.M{}).All(&books)
	return
}

func (r bookRepo) Create(book *schema.Book) (err error) {
	if book.Id.Hex() == "" {
		book.Id = bson.NewObjectId()
	}
	if book.Created.IsZero() {
		book.Created = time.Now()
	}
	book.Updated = time.Now()
	_, err = r.Collection.UpsertId(book.Id, book)
	return
}

func (r bookRepo) Update(book *schema.Book) (err error) {
	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": bson.M{
				"u":    time.Now(),
				"data": book.Data,
				"meta": book.Metadata,
			}}}
	_, err = r.Collection.FindId(book.Id).Apply(change, book)

	return
}
func (r bookRepo) Destroy(id string) (err error) {
	bid := bson.ObjectIdHex(id)
	err = r.Collection.RemoveId(bid)
	return
}

func (r bookRepo) Complete(id string) (book schema.Book, err error) {
	bid := bson.ObjectIdHex(id)
	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": bson.M{
				"cp": time.Now(),
			}}}
	_, err = r.Collection.FindId(bid).Apply(change, &book)

	return
}

func (r bookRepo) Uncomplete(id string) (book schema.Book, err error) {
	bid := bson.ObjectIdHex(id)
	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$unset": bson.M{
				"cp": 1,
			}}}
	_, err = r.Collection.FindId(bid).Apply(change, &book)

	return
}

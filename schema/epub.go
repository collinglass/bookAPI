package schema

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type (
	Books []Book

	Book struct {
		Id       bson.ObjectId `json:"id"           bson:"_id"`
		Metadata Metadata      `json:"meta" bson:"meta"`
		Data     Data          `json:"data" bson:"data"`
		Created  time.Time     `json:"c"            bson:"c"`
		Updated  time.Time     `json:"u,omitempty"  bson:"u,omitempty"`
	}

	Metadata struct {
		Title       []string `json:"title" bson:"t"`
		Language    []string `json:"lang,omitempty" bson:"lang,omitempty"`
		Identifier  []string `json:"id,omitempty" bson:"id,omitempty"`
		Creator     []string `json:"creator,omitempty" bson:"creator,omitempty"`
		Subject     []string `json:"subject,omitempty" bson:"subject,omitempty"`
		Description []string `json:"description,omitempty" bson:"desc,omitempty"`
		Publisher   []string `json:"publisher,omitempty" bson:"pub,omitempty"`
		Contributor []string `json:"contributor,omitempty" bson:"cont,omitempty"`
		Date        []string `json:"date,omitempty" bson:"date,omitempty"`
		EpubType    []string `json:"epubtype,omitempty" bson:"et,omitempty"`
		Format      []string `json:"format,omitempty" bson:"f,omitempty"`
		Source      []string `json:"src,omitempty" bson:"src,omitempty"`
		Relation    []string `json:"rel,omitempty" bson:"rel,omitempty"`
		Coverage    []string `json:"cov,omitempty" bson:"cov,omitempty"`
		Rights      []string `json:"rights,omitempty" bson:"rights,omitempty"`
		Meta        []string `json:"meta,omitempty" bson:"meta,omitempty"`
	}

	Data struct {
		Chapter []Chapter `json:"chapter,omitempty" bson:"chap,omitempty"`
	}

	Chapter struct {
		Title []string `json:"title" bson:"t"`
		Text  []string `json:"text,omitempty" bson:"txt,omitempty"`
	}

	Section struct {
		Title string   `json:"title,omitempty" bson:"t,omitempty"`
		Text  []string `json:"text,omitempty" bson:"txt,omitempty"`
	}
)

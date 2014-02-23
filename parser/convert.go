package main

import (
	"fmt"
	"git.gitorious.org/go-pkg/epubgo.git"
	"log"
)

type Epub struct {
	Metadata struct {
		title       []string
		language    []string
		identifier  []string
		creator     []string
		subject     []string
		description []string
		publisher   []string
		contributor []string
		date        []string
		epubType    []string
		format      []string
		source      []string
		relation    []string
		coverage    []string
		rights      []string
		meta        []string
	}
}

func ExtractMetadata(file string) (Epub, error) {
	// temporary Epub struct
	var temp Epub

	// open epub
	book, err := epubgo.Open(file)

	// defer close until end of func
	defer book.Close()

	// Extract Metadata
	temp.Metadata.title, _ = book.Metadata("title")
	temp.Metadata.language, _ = book.Metadata("language")
	temp.Metadata.identifier, _ = book.Metadata("identifier")
	temp.Metadata.creator, _ = book.Metadata("creator")
	temp.Metadata.subject, _ = book.Metadata("subject")
	temp.Metadata.description, _ = book.Metadata("description")
	temp.Metadata.publisher, _ = book.Metadata("publisher")
	temp.Metadata.contributor, _ = book.Metadata("contributor")
	temp.Metadata.date, _ = book.Metadata("date")
	temp.Metadata.epubType, _ = book.Metadata("type")
	temp.Metadata.format, _ = book.Metadata("format")
	temp.Metadata.source, _ = book.Metadata("source")
	temp.Metadata.relation, _ = book.Metadata("relation")
	temp.Metadata.coverage, _ = book.Metadata("coverage")
	temp.Metadata.rights, _ = book.Metadata("rights")
	temp.Metadata.meta, _ = book.Metadata("meta")

	// Extract other data

	return temp, err
}

func GetMetadata(file Epub) Epub.Metadata {
	// Return file Metadata
	return file.Metadata
}

func ExtractData(file string) (Epub, error) {
	// temporary Epub struct
	var temp Epub

	// open epub
	book, err := epubgo.Open(file)

	// defer close until end of func
	defer book.Close()

	return temp, err
}

func main() {

	test, err := ExtractData("test.epub")

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(test.Metadata)
}

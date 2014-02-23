package main

import (
	"bufio"
	"fmt"
	"git.gitorious.org/go-pkg/epubgo.git"
	"io"
	"log"
	"os"
)

type Epub struct {
	metadata Metadata
	data     *Data
}

type Metadata struct {
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

type Data struct {
	chapter []Chapter
}

type Chapter struct {
	title   string
	section []Section
}

type Section struct {
	title string
	text  []string
}

func ExtractMetadata(file string) (Epub, error) {
	// temporary Epub struct
	var temp Epub

	// open epub
	book, err := epubgo.Open(file)

	// defer close until end of func
	defer book.Close()

	// Extract Metadata
	temp.metadata.title, _ = book.Metadata("title")
	temp.metadata.language, _ = book.Metadata("language")
	temp.metadata.identifier, _ = book.Metadata("identifier")
	temp.metadata.creator, _ = book.Metadata("creator")
	temp.metadata.subject, _ = book.Metadata("subject")
	temp.metadata.description, _ = book.Metadata("description")
	temp.metadata.publisher, _ = book.Metadata("publisher")
	temp.metadata.contributor, _ = book.Metadata("contributor")
	temp.metadata.date, _ = book.Metadata("date")
	temp.metadata.epubType, _ = book.Metadata("type")
	temp.metadata.format, _ = book.Metadata("format")
	temp.metadata.source, _ = book.Metadata("source")
	temp.metadata.relation, _ = book.Metadata("relation")
	temp.metadata.coverage, _ = book.Metadata("coverage")
	temp.metadata.rights, _ = book.Metadata("rights")
	temp.metadata.meta, _ = book.Metadata("meta")

	return temp, err
}

func GetMetadata(file Epub) interface{} {
	// Return file Metadata
	return file.metadata
}

func ExtractData(file string) (Epub, error) {
	// temporary Epub struct

	temp := &Epub{
		metadata: Metadata{
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
			make([]string, 1),
		},
		data: &Data{make([]Chapter, 1)},
	}

	// Initialize inner chapter
	chapter := &Chapter{"", make([]Section, 1, 15)}
	temp.data.chapter[0] = *chapter

	// Initialize inner section
	section := &Section{"", make([]string, 1, 10)}
	temp.data.chapter[0].section[0] = *section

	// Create function to grow slice if not big enough

	// open epub
	book, err := epubgo.Open(file)

	// defer close until end of func
	defer book.Close()

	// Extract data

	/* To extract data we must perform an preorder
	traversal and create a new interface for every new data set */

	// Create Navigation Iterator
	naviter, err := book.Navigation()
	if err != nil {
		return *temp, err
	}

	/* Print all the titles using preorder traversal variant */
	//temp.Data.Chapter = make([]Chapter struct, 10)
	temp.data.chapter[0].title = naviter.Title()
	naviter.In()
	log.Print(naviter.Title())

	for !naviter.IsLast() {
		naviter.Next()
		log.Print(naviter.Title())
		if naviter.HasChildren() {
			naviter.In()
			log.Print(naviter.Title())
			for !naviter.IsLast() {
				naviter.Next()
				log.Print(naviter.Title())
			}
			log.Print(naviter.Title())
			naviter.Out()
		}
	}

	return *temp, err
}

func ReadData() {
	// open input file
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(fi)

	// open output file
	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err = w.Flush(); err != nil {
		panic(err)
	}
}

func main() {

	test, err := ExtractData("test.epub")

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(test.metadata)
}

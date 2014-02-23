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
	Data struct {
		Chapter []struct {
			Section []struct {
				text []string
			}
		}
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

	return temp, err
}

func GetMetadata(file Epub) interface{} {
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

	// Extract data

	/* To extract data we must perform an preorder
	traversal and create a new interface for every new data set */

	// Create Navigation Iterator
	naviter, err := book.Navigation()
	if err != nil {
		return temp, err
	}

	/* Print all the titles using preorder traversal variant */

	log.Print(naviter.Title())
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

	return temp, err
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

	fmt.Println(test.Metadata)
}

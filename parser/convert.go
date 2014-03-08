package main

import (
	"bufio"
	"code.google.com/p/go.net/html"
	"fmt"
	"git.gitorious.org/go-pkg/epubgo.git"
	"io"
	"log"
	"os"
)

type Epub struct {
	metadata Metadata
	data     Data
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

func ExtractData(file string) (Epub, error) {
	temp := initEpub()

	// open epub
	book, err := epubgo.Open(file)
	if err != nil {
		panic(err)
	}

	// defer close until end of func
	defer book.Close()

	//fmt.Println(book.MetadataFields())

	// Create iterator
	it, err := book.Spine()
	if err != nil {
		panic(err)
	}
	it.Next()
	it.Next()
	page, err := it.Open()
	if err != nil {
		panic(err)
	}

	defer page.Close()

	// parse page
	parseHtml(page, temp)

	fmt.Println(*temp)

	return *temp, err
}

func parseHtml(r io.Reader, epub *Epub) {
	d := html.NewTokenizer(r)
	isChap := false
	for {
		// token type
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			return
		}
		token := d.Token()
		switch tokenType {
		case html.StartTagToken:
			if token.Data == "h1" {
				isChap = true
			}
		case html.TextToken:
			if isChap == true {
				chap := &Chapter{token.Data, make([]Section, 1)}
				epub.data.chapter = append(epub.data.chapter, *chap)
				//fmt.Println(*epub)
			}
		case html.EndTagToken:
			if token.Data == "h1" {
				isChap = false
			}
		case html.SelfClosingTagToken: // <tag/>

		}
	}
	return
}

func main() {

	_, err := ExtractData("fingerprint.epub")

	if err != nil {
		log.Panic(err)
	}
}

// ************************************************************
// ************************************************************
// ************************************************************
// ************************************************************

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

func initEpub() *Epub {
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
		data: Data{},
	}

	//initialize data
	data := &Data{make([]Chapter, 1)}
	temp.data = *data
	// Initialize inner chapter
	chapter := &Chapter{"", make([]Section, 1, 15)}
	temp.data.chapter[0] = *chapter

	// Initialize inner section
	section := &Section{"", make([]string, 1, 10)}
	temp.data.chapter[0].section[0] = *section

	return temp
}

func printIndex(book *epubgo.Epub) error {

	// Create Navigation Iterator
	naviter, err := book.Navigation()
	if err != nil {
		return err
	}

	/* Print all the titles using preorder traversal variant */
	log.Println(naviter.Title())
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
			naviter.Out()
		}
	}
	return err
}

func ReadData(file string) error {
	// open input file
	//temp := initEpub()

	// open epub
	book, err := epubgo.Open(file)
	if err != nil {
		panic(err)
	}
	// defer close until end of func
	defer book.Close()

	it, err := book.Spine()
	if err != nil {
		panic(err)
	}
	it.Next()
	it.Next()
	page, err := it.Open()
	if err != nil {
		panic(err)
	}

	defer page.Close()

	// make a read buffer
	r := bufio.NewReader(page)

	// open output file
	fo, err := os.Create("output3.txt")
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
	return err
}

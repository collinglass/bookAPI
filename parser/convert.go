/* Begin convert.go */
package parser

import (
	"bufio"
	"code.google.com/p/go.net/html"
	"encoding/json"
	"errors"
	"fmt"
	"git.gitorious.org/go-pkg/epubgo.git"
	"github.com/collinglass/bookAPI/schema"
	"github.com/collinglass/bookAPI/server/db"
	"io"
	"log"
	"os"
)

func initEpub() *schema.Book {

	// temporary Epub struct
	temp := &schema.Book{
		Metadata: schema.Metadata{
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
		Data: schema.Data{},
	}

	//initialize data
	Data := &schema.Data{make([]schema.Chapter, 1)}
	temp.Data = *Data
	// Initialize inner chapter
	Chapter := &schema.Chapter{"", make([]string, 1, 20)}
	temp.Data.Chapter[0] = *Chapter

	return temp
}

func ExtractMetadata(file string, temp *schema.Book) error {

	// open epub
	book, err := epubgo.Open(file)

	// defer close until end of func
	defer book.Close()

	// Extract Metadata
	(*temp).Metadata.Title, _ = book.Metadata("title")
	(*temp).Metadata.Language, _ = book.Metadata("language")
	(*temp).Metadata.Identifier, _ = book.Metadata("identifier")
	(*temp).Metadata.Creator, _ = book.Metadata("creator")
	(*temp).Metadata.Subject, _ = book.Metadata("subject")
	(*temp).Metadata.Description, _ = book.Metadata("description")
	(*temp).Metadata.Publisher, _ = book.Metadata("publisher")
	(*temp).Metadata.Contributor, _ = book.Metadata("contributor")
	(*temp).Metadata.Date, _ = book.Metadata("date")
	(*temp).Metadata.EpubType, _ = book.Metadata("type")
	(*temp).Metadata.Format, _ = book.Metadata("format")
	(*temp).Metadata.Source, _ = book.Metadata("source")
	(*temp).Metadata.Relation, _ = book.Metadata("relation")
	(*temp).Metadata.Coverage, _ = book.Metadata("coverage")
	(*temp).Metadata.Rights, _ = book.Metadata("rights")
	(*temp).Metadata.Meta, _ = book.Metadata("meta")

	return err
}

func ExtractData(file string) error {
	// Initialize temporary epub to be returned
	temp := initEpub()

	// open epub
	book, err := epubgo.Open(file)

	// defer close until end of func
	defer book.Close()

	// Create iterator
	it, err := book.Spine()

	// Go to page with content
	it.Next()
	it.Next()

	// Open page
	page, err := it.Open()

	// Defer page close til aftr done
	defer page.Close()

	// Extract Metadata
	ExtractMetadata(file, temp)

	// parse page
	parseXHTML(page, temp)

	// Putting in Mongo
	db.InsertMongo(temp)

	return err
}

func createFile(temp *(schema.Book)) error {

	// Epub object for JSON Marshal
	epub := *temp

	fmt.Println(epub)

	// JSON Marshal
	jason, err := json.Marshal(epub)

	// open output file
	fo, err := os.Create("./server/api/v0.1/books/output.json")

	// close fo on exit and check for its returned error
	defer func() {
		err = fo.Close()
	}()

	// Notify user
	log.Println("JSONifying...")

	// make a write buffer
	w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// chunk
		chunk := 1024

		// jason bytes is empty, break
		if len(jason) == 0 {
			break
		}

		// if length of jason is less than 1024, set chunk
		if len(jason) < 1024 {
			chunk = len(jason)
		}

		// Remove a chunk
		buf = jason[:chunk]
		jason = jason[chunk:len(jason)]

		// write a chunk
		if _, err := w.Write(buf); err != nil {
			panic(err)
		}
	}

	// Flush writer
	err = w.Flush()

	// Return epub pointer and error
	return err
}

func parseXHTML(r io.Reader, epub *(schema.Book)) error {
	// Get new Tokenizer
	d := html.NewTokenizer(r)

	// Initialize variables
	isChap := false
	isPar := false
	index := 0

	for {
		// get next token
		tokenType := d.Next()
		// check for error token
		if tokenType == html.ErrorToken {
			return errors.New("parseXHTML Failed: Error token when parsing!")
		}
		// Get current token
		token := d.Token()

		switch tokenType {

		// Start token
		case html.StartTagToken:
			if token.Data == "h3" {
				isChap = true
			}
			if token.Data == "p" {
				isPar = true
			}

		// Text token
		case html.TextToken:
			if isChap == true {
				chap := &schema.Chapter{token.Data, make([]string, 1, 20)}
				epub.Data.Chapter = append(epub.Data.Chapter, *chap)
			}
			if isPar == true {
				epub.Data.Chapter[index].Text = append(epub.Data.Chapter[index].Text, token.Data)
			}

		// End token
		case html.EndTagToken:
			if token.Data == "h3" {
				isChap = false
				index++
			}
			if token.Data == "p" {
				isPar = false
			}

		// Self closing token
		case html.SelfClosingTagToken: // <tag/>

		}
	}
	return nil
}

/* End convert.go */

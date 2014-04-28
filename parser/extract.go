/* Begin convert.go */
package parser

import (
	"code.google.com/p/go.net/html"
	"errors"
	"git.gitorious.org/go-pkg/epubgo.git"
	"github.com/collinglass/bookAPI/schema"
	"io"
	"strings"
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
	Chapter := &schema.Chapter{make([]string, 0), make([]string, 0)}
	temp.Data.Chapter[0] = *Chapter

	return temp
}

func extractMetadata(file string, temp *schema.Book) error {

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

func extractData(file string) (*schema.Book, error) {
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
	err = extractMetadata(file, temp)

	// parse page
	parseXHTML(page, temp)

	return temp, err
}

func parseXHTML(r io.Reader, epub *schema.Book) error {
	// Get new Tokenizer
	d := html.NewTokenizer(r)

	// Initialize variables
	isChap := false
	isPar := false
	isTitle := false
	index := 0
	isEmpty := false

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
			if token.Data == "h4" {
				isTitle = true
			}
			if token.Data == "p" {
				isPar = true
			}

		// Text token
		case html.TextToken:
			if len(strings.Fields(token.Data)) == 0 {
				isEmpty = true
			}
			if isChap && !isEmpty {
				chap := &schema.Chapter{make([]string, 0), make([]string, 0)}
				chap.Title = append(chap.Title, token.Data)
				epub.Data.Chapter = append(epub.Data.Chapter, *chap)
			}
			if isTitle && !isEmpty {
				epub.Data.Chapter[index].Title = append(epub.Data.Chapter[index].Title, token.Data)
			}
			if isPar && !isEmpty {
				epub.Data.Chapter[index].Text = append(epub.Data.Chapter[index].Text, token.Data)
			}
			isEmpty = false

		// End token
		case html.EndTagToken:
			if token.Data == "h3" {
				isChap = false
				index++
			}
			if token.Data == "h4" {
				isTitle = false
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

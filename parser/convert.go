package parser

import (
	"bufio"
	"code.google.com/p/go.net/html"
	"encoding/json"
	//"fmt"
	"git.gitorious.org/go-pkg/epubgo.git"
	"io"
	"io/ioutil"
	//"log"
	"os"
)

type Epub struct {
	Metadata Metadata
	Data     Data
}

type Metadata struct {
	Title       []string
	Language    []string
	Identifier  []string
	Creator     []string
	Subject     []string
	Description []string
	Publisher   []string
	Contributor []string
	Date        []string
	EpubType    []string
	Format      []string
	Source      []string
	Relation    []string
	Coverage    []string
	Rights      []string
	Meta        []string
}

type Data struct {
	Chapter []Chapter
}

type Chapter struct {
	Title   string
	Section []Section
	Text    []string
}

type Section struct {
	Title string
	Text  []string
}

func initEpub() *Epub {

	// temporary Epub struct
	temp := &Epub{
		Metadata: Metadata{
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
		Data: Data{},
	}

	//initialize data
	Data := &Data{make([]Chapter, 1)}
	temp.Data = *Data
	// Initialize inner chapter
	Chapter := &Chapter{"", make([]Section, 1, 15), make([]string, 1, 15)}
	temp.Data.Chapter[0] = *Chapter

	// Initialize inner section
	Section := &Section{"", make([]string, 1, 10)}
	temp.Data.Chapter[0].Section[0] = *Section

	return temp
}

func ExtractMetadata(file string, temp *Epub) error {

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

func ExtractData(file string) (Epub, error) {
	// Initialize temporary epub to be returned
	temp := initEpub()

	// open epub
	book, err := epubgo.Open(file)
	if err != nil {
		panic(err)
	}

	// defer close until end of func
	defer book.Close()

	// Create iterator
	it, err := book.Spine()
	if err != nil {
		panic(err)
	}
	// Go to page with content
	it.Next()
	it.Next()

	// Open page
	page, err := it.Open()
	if err != nil {
		panic(err)
	}

	// Defer page close til aftr done
	defer page.Close()

	// Extract Metadata
	ExtractMetadata(file, temp)

	// parse page
	parseXHTML(page, temp)

	// Epub object for JSON Marshal
	epub := *temp

	// JSON Marshal
	jason, err := json.Marshal(epub)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./server/api/v0.1/books/output.json", jason, 0644)

	return *temp, err
}

func parseXHTML(r io.Reader, epub *Epub) {
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
			return
		}
		// Get current token
		token := d.Token()

		switch tokenType {

		// Start token
		case html.StartTagToken:
			if token.Data == "h1" {
				isChap = true
			}
			if token.Data == "p" {
				isPar = true
			}

		// Text token
		case html.TextToken:
			if isChap == true {
				chap := &Chapter{token.Data, make([]Section, 1), make([]string, 1, 15)}
				epub.Data.Chapter = append(epub.Data.Chapter, *chap)
			}
			if isPar == true {
				epub.Data.Chapter[index].Text = append(epub.Data.Chapter[index].Text, token.Data)
			}

		// End token
		case html.EndTagToken:
			if token.Data == "h1" {
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
	return
}

// ************************************************************
// ************************************************************
// ************************************************************
// ************************************************************

func ReadData(file string) error {

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

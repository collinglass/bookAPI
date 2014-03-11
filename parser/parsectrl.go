/* Begin parsectrl.go */
package parser

import (
	"encoding/json"
	"fmt"
	"github.com/collinglass/bookAPI/server/db"
	"log"
)

// Control function to insert book into Database
func InsertBook(file string) error {

	var err error

	// Extract Data
	book, err := extractData(file)

	// Putting in Mongo
	db.InsertMongo(book)

	return err
}

// Control Function to create XHTML file
func ConvertToXHTML(file string) error {

	//epub, err := openEpub(file)

	err := createXHTMLFile(file, "./textcheck/output.xhtml")

	return err
}

// Control Function to create JSON file
func ConvertToJSON(file string) error {

	// Epub object for JSON Marshal
	epub, err := extractData(file)

	// Notify user
	log.Println("JSONifying...")

	// Print epub
	fmt.Println(epub)

	// JSON Marshal
	jason, err := json.Marshal(epub)

	// Create ouputFile
	createFile(&jason, "./server/api/v0.1/books/output.json")

	return err
}

/* End parsectrl.go */

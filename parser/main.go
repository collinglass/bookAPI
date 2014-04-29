package parser

import (
	"bytes"
	"flag"
	"fmt"
)

func main() {

	var err error

	wordPtr := flag.String("word", "15", "a string")

	flag.Parse()

	fmt.Println("word:", *wordPtr)

	var buffer bytes.Buffer

	buffer.WriteString("textcheck/")
	buffer.WriteString("MVP/")
	buffer.WriteString(*wordPtr)
	buffer.WriteString(".epub")

	//err = parser.InsertBook(buffer.String())

	err = ConvertToXHTML(buffer.String())

	err = ConvertToJSON(buffer.String())
	if err != nil {
		panic(err)
	}
}

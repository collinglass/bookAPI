package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/collinglass/bookAPI/parser"
)

func main() {

	var err error

	wordPtr := flag.String("word", "foo", "a string")

	flag.Parse()

	fmt.Println("word:", *wordPtr)

	var buffer bytes.Buffer

	buffer.WriteString("textcheck/")
	buffer.WriteString(*wordPtr)
	buffer.WriteString(".epub")
	/*
		err := parser.ConvertData(buffer.String())
		if err != nil {
			panic(err)
		}
	*/
	err = parser.ExtractData(buffer.String())
	if err != nil {
		panic(err)
	}
}

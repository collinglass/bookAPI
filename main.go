package main

import (
	"github.com/collinglass/bookAPI/parser"
)

func main() {

	_, err := parser.ExtractData("parser/fingerprint.epub")
	if err != nil {
		panic(err)
	}
}

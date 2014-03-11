package parser

import (
	"bufio"
	"git.gitorious.org/go-pkg/epubgo.git"
	"io"
	"os"
)

// Create new file
func createXHTMLFile(input string, output string) error {
	// open epub
	book, err := epubgo.Open(input)
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
	fo, err := os.Create(output)
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

func createFile(inputslice *[]byte, file string) error {
	// Get slice
	bytes := *inputslice

	// open output file
	fo, err := os.Create(file)

	// close fo on exit and check for its returned error
	defer func() {
		err = fo.Close()
	}()

	// make a write buffer
	w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// chunk
		chunk := 1024

		// bytes is empty, break
		if len(bytes) == 0 {
			break
		}

		// if length of bytes is less than 1024, set chunk
		if len(bytes) < 1024 {
			chunk = len(bytes)
		}

		// Remove a chunk
		buf = bytes[:chunk]
		bytes = bytes[chunk:len(bytes)]

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

// Open epub
func openEpub(file string) (*io.ReadCloser, error) {

	// open epub
	book, err := epubgo.Open(file)

	// defer close until end of func
	defer book.Close()

	it, err := book.Spine()

	it.Next()
	it.Next()
	page, err := it.Open()

	return &page, err
}

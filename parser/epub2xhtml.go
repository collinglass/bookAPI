package parser

import (
	"bufio"
	"git.gitorious.org/go-pkg/epubgo.git"
	"io"
	"os"
)

func ConvertData(file string) error {

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
	fo, err := os.Create("./textcheck/output.xhtml")
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

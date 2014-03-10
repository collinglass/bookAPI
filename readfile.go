package main

import (
	"bufio"
	"git.gitorious.org/go-pkg/epubgo.git"
	"io"
	"log"
	"os"
)

func main() {

	err := ReadData("fingerprint.epub")

	if err != nil {
		log.Panic(err)
	}
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

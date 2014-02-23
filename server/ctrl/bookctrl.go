package ctrl

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func BookCtrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}

func GetBookList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filepath := "./api/v0.1/books/books.json"

		fi, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}
		// close fi on exit and check for its returned error
		defer func() {
			if err := fi.Close(); err != nil {
				panic(err)
			}
		}()

		// make a buffer to keep chunks that are read
		buf := make([]byte, 1024)
		for {
			// read a chunk
			n, err := fi.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}

			// write a chunk
			if _, err := w.Write([]byte(buf[:n])); err != nil {
				panic(err)
			}
		}
	}
}

func GetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// open input file

		w.Header().Set("Content-Type", "application/json")
		dir, file := path.Split(r.URL.String())
		fmt.Printf("Path: %v File: %v\n", dir, file)

		filepath := "./api/v0.1/books/" + file + ".json"

		fmt.Printf(filepath)

		fi, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}
		// close fi on exit and check for its returned error
		defer func() {
			if err := fi.Close(); err != nil {
				panic(err)
			}
		}()

		// make a buffer to keep chunks that are read
		buf := make([]byte, 1024)
		for {
			// read a chunk
			n, err := fi.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}

			// write a chunk
			if _, err := w.Write([]byte(buf[:n])); err != nil {
				panic(err)
			}
		}
	}
}

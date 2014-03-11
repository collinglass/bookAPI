Static website for Proof of Concept of BookAPI

# Usage

Read the Docs page

# Development

Install Go with proper [structure](http://golang.org/doc/code.html)

```git clone``` this repository

Go to bookAPI/server

```go run server.go```

# New books

Follow development but stay in main directory

In your text editor load bookAPI/main.go

Change the file path to the location of your epub to convert.

```go run main.go```

File is now located in bookAPI/books/file.json

## TODO - Front-End

Create A-Z page

Create Docs page

Create Tutorial page

## TODO - Back-End

#### General

Remove all redundancies (not yet)

Remove all .epub, .txt files (not yet)

#### Parser

Convert 12 books into json (2/12) (easy)

Abstract Epub Extraction (HARD)

Abstract ToJSON/ToXHTML (easy-Medium)

#### DB

Retrieve books from mongoDB to serve to API endpoint (Medium)

Create api endpoints for all books, books by category (Medium)

#### Schema

Abstract epub schema for parser (Medium)









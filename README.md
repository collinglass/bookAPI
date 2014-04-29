Static website for Proof of Concept of BookAPI

# Usage

Read the Docs page

# Current Version

Book list: ```/api/books/```

15 books: ```/api/books/{1-15}```

1. The Fingerprint
2. All Else is Folly
3. Lonely Road
4. The Brading Collection
5. The Last Battle
6. The Horse and His Boy
7. The Silver Chair
8. Prince Caspian
9. The Voyage of the Dawn Treader
10. Le Major Anspech
11. Napoléon
12. Comment s'est faite la Restauration de 1814
13. L'Horloge qui chante
14. The Last Laird of MacNab
15. La Nation canadienne

# Development

1. Install Go with proper [structure](http://golang.org/doc/code.html)
2. ```git clone``` this repository
3. Go to bookAPI/server
4. ```go run server.go```
5. Go to ```localhost:3000```

# New books

1. Follow development but stay in main directory
2. In your text editor load ```/main.go```
3. Change the file path to the location of your epub to convert.
4. ```go run main.go```
	An XHTML file is now in ```textcheck/``` and a JSON file is in ```server/api/v0.1/books/```
5. Tweak the ```parser/extract.go``` file to display the correct JSON.
	Repeat steps 4 and 5 til right.
6. Comment the XHTML, JSON parsing and Uncomment the ```db.InsertBook()``` line in ```main.go```
7. Execute 4.

## TODO - Front-End

Update landing page

## TODO - Back-End

#### Parser

Write Bash Script to bulk parse all epubs

#### DB

<b>MongoDB put on hold:</b> serving raw JSON files to provide easier way to update ugly JSON!

Retrieve books from mongoDB to serve to API endpoint (Medium)

API endpoints to Get from Mongo

#### Schema

Abstract epub schema for parser (Medium)

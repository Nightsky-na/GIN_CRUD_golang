# Library Management API with Gin Framework

This program is a simple RESTful CRUD (Create, Read, Update, Delete) API for a library management system using the Go programming language and the Gin web framework. Gin is a high-performance HTTP web framework that makes it easy to build robust APIs. It demonstrates how to build an API server with endpoints for managing a collection of books. The program uses the Gin package to handle HTTP routing, middleware, and JSON encoding/decoding.

The source code can be found at:  [URL GitHub](https://github.com/techwithtim/Go-API-Tutorial/tree/main)

## Features
* Create a new book
* Retrieve all books
* Retrieve a single book by ID
* Update a book by ID
* Delete a book by ID
* Check out a book
* Return a book

## Dependencies
To run this program, you will need to install the Gin web framework package:
```
go get -u github.com/gin-gonic/gin
```

## Data Structures
The program uses two main data structures:

* `book`: Represents a book with the following fields:
    * ID: Unique identifier for the book
    * Title: Title of the book
    * Author: Author of the book
    * Quantity: Number of copies of the book in the library

## API Endpoints
The API has the following endpoints:

* `GET /books`: Retrieve all books
* `GET /books/:id`: Retrieve a single book by ID
* `PUT /books`: Create a new book
* `PATCH /checkout`: Check out a book
* `PATCH /return`: Return a book
* `DELETE /books/:id`: Delete a book by ID

## Usage
To run the program, execute the following command:
```
go run main.go
```

The API will start listening on port 8080. The Gin framework provides efficient request handling, easy middleware integration, and graceful error handling, making the development process smooth and enjoyable.

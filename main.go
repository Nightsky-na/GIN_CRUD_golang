package main

// Download gin =>  go get -u github.com/gin-gonic/gin
import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	//capital letter for public
	ID       string `json: "id"`
	Title    string `json: "title"`
	Author   string `json: "author"`
	Quantity int    `json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// return all books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// Create a new book
/*
curl localhost:8080/books --include --header "Content-Type: application/json" --request PUT --data @body.json
*/
func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		// send bad request if the JSON is not valid
		return
	}

	// add the new book to the slice
	books = append(books, newBook)

	// return the new book to the client
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := bookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// get book by id
func bookByID(id string) (*book, error) {
	for index, book := range books {
		if book.ID == id {
			return &books[index], nil
		}
	}
	return nil, errors.New("book not found")
}

// check out book
//curl "localhost:8080/checkout?id=1" --request PATCH

func checkOutBook(c *gin.Context) {
	// get id from query string
	id, ok := c.GetQuery("id")

	// if id is not provided
	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	// get book by id
	book, err := bookByID(id)

	// if book is not found
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	// if book is not available
	if book.Quantity == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book is not available"})
		return
	}

	book.Quantity = book.Quantity - 1
	c.IndentedJSON(http.StatusOK, book)

}

// return book
// curl "localhost:8080/return?id=1" --request PATCH
func returnBook(c *gin.Context) {
	// get id from query string
	id, ok := c.GetQuery("id")

	// if id is not provided
	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	// get book by id
	book, err := bookByID(id)

	// if book is not found
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity = book.Quantity + 1
	c.IndentedJSON(http.StatusOK, book)

}

func main() {
	// setup router
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.PUT("/books", createBook)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/return", returnBook)

	router.Run("localhost:8080")
}

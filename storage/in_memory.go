package storage

import (
	"sync"

	"github.com/google/uuid"
	"github.com/Shashika2001/bookstore-api/models"
)

var (
	books = make(map[string]models.Book)
	mu    sync.Mutex
)

// GetAllBooks returns all books
func GetAllBooks() []models.Book {
	mu.Lock()
	defer mu.Unlock()
	var bookList []models.Book
	for _, book := range books {
		bookList = append(bookList, book)
	}
	return bookList
}

// GetBookByID retrieves a book by ID
func GetBookByID(id string) (models.Book, bool) {
	mu.Lock()
	defer mu.Unlock()
	book, exists := books[id]
	return book, exists
}

// CreateBook adds a new book
func CreateBook(book models.Book) models.Book {
	mu.Lock()
	defer mu.Unlock()
	book.BookId = uuid.New().String()
	books[book.BookId] = book
	return book
}

// UpdateBook updates an existing book
func UpdateBook(id string, updatedBook models.Book) (models.Book, bool) {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := books[id]; !exists {
		return models.Book{}, false
	}
	updatedBook.BookId = id
	books[id] = updatedBook
	return updatedBook, true
}

// DeleteBook removes a book by ID
func DeleteBook(id string) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := books[id]; exists {
		delete(books, id)
		return true
	}
	return false
}

package storage

import (
	"strings"
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

func SearchBooksConcurrently(query string) []models.Book {
	mu.Lock()
	defer mu.Unlock()

	numWorkers := 4 // Number of goroutines
	bookList := GetAllBooks()
	bookCount := len(bookList)
	if bookCount == 0 {
		return []models.Book{}
	}

	// Define chunk size for each worker
	chunkSize := (bookCount + numWorkers - 1) / numWorkers

	// Channel to collect results
	results := make(chan []models.Book, numWorkers)
	var wg sync.WaitGroup

	// Split search into concurrent tasks
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > bookCount {
			end = bookCount
		}
		wg.Add(1)
		go func(subBooks []models.Book) {
			defer wg.Done()
			matchingBooks := searchBooks(subBooks, query)
			results <- matchingBooks
		}(bookList[start:end])
	}

	// Close results channel once all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results from all workers
	var finalResults []models.Book
	for books := range results {
		finalResults = append(finalResults, books...)
	}

	return finalResults
}

// searchBooks filters books by checking if the query is present in title or description
func searchBooks(books []models.Book, query string) []models.Book {
	var matched []models.Book
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), query) ||
			strings.Contains(strings.ToLower(book.Description), query) {
			matched = append(matched, book)
		}
	}
	return matched
}

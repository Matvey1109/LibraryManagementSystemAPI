package repositories

import (
	"app/internal/models"
	"log"
)

type BookRepository struct{}

func (br *BookRepository) GetAllBooks() []models.Book {
	books, err := storage.GetAllBooksStorage()
	if err != nil {
		log.Fatal(err)
	}
	return books
}

func (br *BookRepository) GetBook(id string) models.Book {
	book, err := storage.GetBookStorage(id)
	if err != nil {
		log.Fatal(err)
	}
	return book
}

func (br *BookRepository) AddBook(title string, author string, publicationYear int, genre string, totalCopies int) {
	newID := GenerateID()
	newBook := models.Book{
		ID:              newID,
		Title:           title,
		Author:          author,
		PublicationYear: publicationYear,
		Genre:           genre,
		AvailableCopies: totalCopies,
		TotalCopies:     totalCopies,
	}
	err := storage.AddBookStorage(newBook)
	if err != nil {
		log.Fatal(err)
	}
}

func (br *BookRepository) UpdateBook(id string, title string, author string, publicationYear int, genre string, availableCopies int, totalCopies int) {
	book, err := storage.GetBookStorage(id)
	if err != nil {
		log.Fatal(err)
	}

	if availableCopies > totalCopies {
		log.Fatal("available copies cannot be greater than total copies")
	}

	if title != "" {
		book.Title = title
	}
	if author != "" {
		book.Author = author
	}
	if publicationYear != -1 {
		book.PublicationYear = publicationYear
	}
	if genre != "" {
		book.Genre = genre
	}
	if availableCopies >= 0 {
		book.AvailableCopies = availableCopies
	}
	if totalCopies >= 0 {
		book.TotalCopies = totalCopies
	}

	err = storage.UpdateBookStorage(id, book)
	if err != nil {
		log.Fatal(err)
	}
}

func (br *BookRepository) DeleteBook(id string) {
	err := storage.DeleteBookStorage(id)
	if err != nil {
		log.Fatal(err)
	}
}

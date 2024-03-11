package repositories

import (
	"app/internal/models"
	"errors"
)

type BookRepository struct{}

func (br *BookRepository) GetAllBooks() ([]models.Book, error) {
	books, err := storage.GetAllBooksStorage()
	if err != nil {
		return books, err
	}
	return books, nil
}

func (br *BookRepository) GetBook(id string) (models.Book, error) {
	book, err := storage.GetBookStorage(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (br *BookRepository) AddBook(title string, author string, publicationYear int, genre string, totalCopies int) error {
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
		return err
	}
	return nil
}

func (br *BookRepository) UpdateBook(id string, title string, author string, publicationYear int, genre string, availableCopies int, totalCopies int) error {
	book, err := storage.GetBookStorage(id)
	if err != nil {
		return err
	}

	if availableCopies > totalCopies {
		return errors.New("available copies cannot be greater than total copies")
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
		return err
	}
	return err
}

func (br *BookRepository) DeleteBook(id string) error {
	err := storage.DeleteBookStorage(id)
	if err != nil {
		return err
	}
	return nil
}

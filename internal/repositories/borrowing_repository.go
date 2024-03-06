package repositories

import (
	"app/internal/models"
	"log"
	"time"
)

type BorrowingRepository struct{}

func (br *BorrowingRepository) GetAllBorrowings() []models.Borrowing {
	borrowings, err := storage.GetAllBorrowingsStorage()
	if err != nil {
		log.Fatal(err)
	}
	return borrowings
}

func (br *BorrowingRepository) GetMemberBooks(memberID string) []models.Book {
	borrowings, err := storage.GetAllBorrowingsStorage()
	if err != nil {
		log.Fatal(err)
	}

	memberBorrowings := make([]models.Borrowing, 0)
	for _, borrowing := range borrowings {
		if borrowing.MemberID == memberID {
			memberBorrowings = append(memberBorrowings, borrowing)
		}
	}

	books := make([]models.Book, 0)
	for _, borrowing := range memberBorrowings {
		book, err := storage.GetBookStorage(borrowing.BookID)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	return books
}

func (br *BorrowingRepository) BorrowBook(bookID string, memberID string, borrowYear int) {
	newID := GenerateID()
	newBorrowing := models.Borrowing{
		ID:         newID,
		BookID:     bookID,
		MemberID:   memberID,
		BorrowYear: borrowYear,
		ReturnYear: -1,
	}
	err := storage.AddBorrowingStorage(newBorrowing)
	if err != nil {
		log.Fatal(err)
	}
}

func (br *BorrowingRepository) ReturnBook(id string) {
	borrowing, err := storage.GetBorrowingStorage(id)
	if err != nil {
		log.Fatal(err)
	}

	borrowing.ReturnYear = time.Now().Year()

	err = storage.UpdateBorrowingStorage(id, borrowing)
	if err != nil {
		log.Fatal(err)
	}
}

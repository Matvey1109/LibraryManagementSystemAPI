package utils

import (
	"github.com/Matvey1109/LibraryManagementSystemAPI/internal/models"
	"github.com/Matvey1109/LibraryManagementSystemSerializers/schemas"
)

func ConvertMemberToSchema(member models.Member) schemas.Member {
	convertedMember := schemas.Member{
		ID:        member.ID,
		Name:      member.Name,
		Address:   member.Address,
		Email:     member.Email,
		CreatedAt: member.CreatedAt,
	}
	return convertedMember
}

func ConvertBookToSchema(book models.Book) schemas.Book {
	convertedBook := schemas.Book{
		ID:              book.ID,
		Title:           book.Title,
		Author:          book.Author,
		PublicationYear: book.PublicationYear,
		Genre:           book.Genre,
		AvailableCopies: book.AvailableCopies,
		TotalCopies:     book.TotalCopies,
	}
	return convertedBook
}

func ConvertBorrowingToSchema(borrowing models.Borrowing) schemas.Borrowing {
	convertedBorrowing := schemas.Borrowing{
		ID:         borrowing.ID,
		BookID:     borrowing.BookID,
		MemberID:   borrowing.MemberID,
		BorrowYear: borrowing.BorrowYear,
		ReturnYear: borrowing.ReturnYear,
	}
	return convertedBorrowing
}

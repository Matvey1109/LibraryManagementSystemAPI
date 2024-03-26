package console

import (
	"app/internal/repositories"
	"fmt"
	"log"
)

func TestMember() {
	memberRepo := repositories.ExportRepository

	err := memberRepo.AddMember("Example member", "any address", "test@gmail.com")
	if err != nil {
		log.Fatal("Failed to add member:", err)
	}
	fmt.Println("Member added successfully.")

	members, err := memberRepo.GetAllMembers()
	if err != nil {
		log.Fatal("Failed to get members:", err)
	}
	var tempId string
	fmt.Println("All Members:")
	for _, member := range members {
		fmt.Printf("ID: %s, Name: %s, Address: %s, Email: %s\n", member.ID, member.Name, member.Address, member.Email)
		tempId = member.GetMemberID()
	}

	newName := "Updated name"
	err = memberRepo.UpdateMember(tempId, &newName, nil, nil)
	if err != nil {
		log.Fatal("Failed to update member:", err)
	}
	fmt.Println("Member updated successfully.")

	members, err = memberRepo.GetAllMembers()
	if err != nil {
		log.Fatal("Failed to get members:", err)
	}
	fmt.Println("All Members:")
	for _, member := range members {
		fmt.Printf("ID: %s, Name: %s, Address: %s, Email: %s\n", member.ID, member.Name, member.Address, member.Email)
		tempId = member.GetMemberID()
	}

	err = memberRepo.DeleteMember(tempId)
	if err != nil {
		log.Fatal("Failed to delete book:", err)
	}
	fmt.Println("Book deleted successfully.")

	members, err = memberRepo.GetAllMembers()
	if err != nil {
		log.Fatal("Failed to get members:", err)
	}
	fmt.Println("All Members:")
	for _, member := range members {
		fmt.Printf("ID: %s, Name: %s, Address: %s, Email: %s\n", member.ID, member.Name, member.Address, member.Email)
	}
	fmt.Println()
}

func TestBook() {
	bookRepo := repositories.ExportRepository

	err := bookRepo.AddBook("Example title", "F. Scott Fitzgerald", 1925, "Fiction", 5)
	if err != nil {
		log.Fatal("Failed to add book:", err)
	}
	fmt.Println("Book added successfully.")

	books, err := bookRepo.GetAllBooks()
	if err != nil {
		log.Fatal("Failed to get books:", err)
	}
	var tempId string
	fmt.Println("All Books:")
	for _, book := range books {
		fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		tempId = book.GetBookID()
	}

	newTitle := "Updated title"
	err = bookRepo.UpdateBook(tempId, &newTitle, nil, nil, nil, nil, nil)
	if err != nil {
		log.Fatal("Failed to update book:", err)
	}
	fmt.Println("Book updated successfully.")

	books, err = bookRepo.GetAllBooks()
	if err != nil {
		log.Fatal("Failed to get books:", err)
	}
	fmt.Println("All Books:")
	for _, book := range books {
		fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}

	err = bookRepo.DeleteBook(tempId)
	if err != nil {
		log.Fatal("Failed to delete book:", err)
	}
	fmt.Println("Book deleted successfully.")

	books, err = bookRepo.GetAllBooks()
	if err != nil {
		log.Fatal("Failed to get books:", err)
	}
	fmt.Println("All Books:")
	for _, book := range books {
		fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
	fmt.Println()
}

func TestBorrowing() {
	borrowingRepo := repositories.ExportRepository

	err := borrowingRepo.BorrowBook("bookID1", "memberID1", 2021)
	if err != nil {
		log.Fatal("Failed to borrow book:", err)
	}
	fmt.Println("Book borrowed successfully.")

	var tempId string
	borrrowings, err := borrowingRepo.GetAllBorrowings()
	if err != nil {
		log.Fatal("Failed to get borrowings:", err)
	}
	fmt.Println("All borrowings:")
	for _, borrowing := range borrrowings {
		fmt.Printf("ID: %s, BookID: %s, MemberID: %s, BorrowYear: %d\n", borrowing.ID, borrowing.BookID, borrowing.MemberID, borrowing.BorrowYear)
		tempId = borrowing.GetBorrowingID()
	}

	err = borrowingRepo.ReturnBook(tempId)
	if err != nil {
		log.Fatal("Failed to return book:", err)
	}
	fmt.Println("Book returned successfully.")

	borrrowings, err = borrowingRepo.GetAllBorrowings()
	if err != nil {
		log.Fatal("Failed to get borrowings:", err)
	}
	fmt.Println("All borrowings:")
	for _, borrowing := range borrrowings {
		fmt.Printf("ID: %s, BookID: %s, MemberID: %s, BorrowYear: %d, ReturnYear: %d\n", borrowing.ID, borrowing.BookID, borrowing.MemberID, borrowing.BorrowYear, borrowing.ReturnYear)
	}
	fmt.Println()
}

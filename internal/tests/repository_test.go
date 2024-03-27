package tests

import (
	"testing"

	"github.com/Matvey1109/LibraryManagementSystemAPI/internal/models"
	"github.com/Matvey1109/LibraryManagementSystemAPI/internal/repositories"

	"github.com/Matvey1109/LibraryManagementSystemSerializers/models"
	"github.com/Matvey1109/LibraryManagementSystemSerializers/serializers/json_serializers"
	"github.com/Matvey1109/LibraryManagementSystemSerializers/serializers/xml_serializers"
)

func TestMemberRepository(t *testing.T) {
	memberRepo := repositories.NewRepository()

	// AddMember
	memberRepo.AddMember("Example member", "any address", "test@gmail.com")
	t.Logf("Calling AddMember()")

	// GetAllMembers
	result, _ := memberRepo.GetAllMembers()
	expected := make([]models.Member, 1)
	t.Logf("Calling GetAllMembers(), result = %v\n", result)
	if len(result) != len(expected) {
		t.Errorf("ERROR. Expected: %v, got: %v", expected, result)
	}

	// SerializeMemberToJson
	jsonMember, _ := json_serializers.SerializeMemberToJson(result[0])
	t.Logf("Calling SerializeMemberToJson(), result = \n%v\n", string(jsonMember))

	// DeserializeMemberFromJson
	member, _ := json_serializers.DeserializeMemberFromJson(jsonMember)
	t.Logf("Calling DeserializeMemberFromJson(), result = %v\n", member)

	// UpdateMember
	tempID := result[0].GetMemberID()
	newName := "Updated name"
	memberRepo.UpdateMember(tempID, &newName, nil, nil)
	t.Logf("Calling UpdateMember()")

	// GetAllMembers
	result, _ = memberRepo.GetAllMembers()
	expected = make([]models.Member, 1)
	t.Logf("Calling GetAllMembers(), result = %v\n", result)
	if len(result) != len(expected) {
		t.Errorf("ERROR. Expected: %v, got: %v", expected, result)
	}

	// SerializeMemberToXML
	XMLMember, _ := xml_serializers.SerializeMemberToXML(result[0])
	t.Logf("Calling SerializeMemberToXML(), result = \n%v\n", string(XMLMember))

	// DeserializeMemberFromXML
	member, _ = xml_serializers.DeserializeMemberFromXML(XMLMember)
	t.Logf("Calling DeserializeMemberFromXML(), result = %v\n", member)

	// DeleteMember
	memberRepo.DeleteMember(tempID)
	t.Logf("Calling DeleteMember()")

	// GetAllMembers
	result, _ = memberRepo.GetAllMembers()
	expected = make([]models.Member, 0)
	t.Logf("Calling GetAllMembers(), result = %v\n", result)
	if len(result) != len(expected) {
		t.Errorf("ERROR. Expected: %v, got: %v", expected, result)
	}
}

func TestBookRepository(t *testing.T) {
	bookRepo := repositories.NewRepository()

	// AddBook
	bookRepo.AddBook("Example title", "F. Scott Fitzgerald", 1925, "Fiction", 5)
	t.Logf("Calling AddBook()")

	// GetAllBooks
	result, _ := bookRepo.GetAllBooks()
	expected := make([]models.Book, 1)
	t.Logf("Calling GetAllBooks(), result = %v\n", result)
	if len(result) != len(expected) {
		t.Errorf("ERROR. Expected: %v, got: %v", expected, result)
	}

	// SerializeBookToJson
	jsonBook, _ := json_serializers.SerializeBookToJson(result[0])
	t.Logf("Calling SerializeBookToJson(), result = \n%v\n", string(jsonBook))

	// DeserializeBookFromJson
	book, _ := json_serializers.DeserializeBookFromJson(jsonBook)
	t.Logf("Calling DeserializeBookFromJson(), result = %v\n", book)

	// UpdateBook
	tempID := result[0].GetBookID()
	newTitle := "Updated title"
	bookRepo.UpdateBook(tempID, &newTitle, nil, nil, nil, nil, nil)
	t.Logf("Calling UpdateBook()")

	// GetAllBooks
	result, _ = bookRepo.GetAllBooks()
	expected = make([]models.Book, 1)
	t.Logf("Calling GetAllBooks(), result = %v\n", result)
	if len(result) != len(expected) {
		t.Errorf("ERROR. Expected: %v, got: %v", expected, result)
	}

	// SerializeBookToXML
	XMLBook, _ := xml_serializers.SerializeBookToXML(result[0])
	t.Logf("Calling SerializeBookToXML(), result = \n%v\n", string(XMLBook))

	// DeserializeBookFromXML
	book, _ = xml_serializers.DeserializeBookFromXML(XMLBook)
	t.Logf("Calling DeserializeBookFromXML(), result = %v\n", book)

	// DeleteBook
	bookRepo.DeleteBook(tempID)
	t.Logf("Calling DeleteBook()")

	// GetAllBooks
	result, _ = bookRepo.GetAllBooks()
	expected = make([]models.Book, 0)
	t.Logf("Calling GetAllBooks(), result = %v\n", result)
	if len(result) != len(expected) {
		t.Errorf("ERROR. Expected: %v, got: %v", expected, result)
	}
}

func TestBorrowingRepository(t *testing.T) {
	borrowingRepo := repositories.NewRepository()

	// BorrowBook
	borrowingRepo.BorrowBook("bookID1", "memberID1", 2021)
	t.Logf("Calling BorrowBook()")

	// GetAllBorrowings
	result, _ := borrowingRepo.GetAllBorrowings()
	expected := make([]models.Borrowing, 1)
	t.Logf("Calling GetAllBorrowings(), result = %v\n", result)
	if len(result) != len(expected) {
		t.Errorf("ERROR. Expected: %v, got: %v", expected, result)
	}

	// SerializeBorrowingToJson
	jsonBorrowing, _ := json_serializers.SerializeBorrowingToJson(result[0])
	t.Logf("Calling SerializeBorrowingToJson(), result = \n%v\n", string(jsonBorrowing))

	// DeserializeBorrowingFromJson
	borrowing, _ := json_serializers.DeserializeBorrowingFromJson(jsonBorrowing)
	t.Logf("Calling DeserializeBorrowingFromJson(), result = %v\n", borrowing)

	// ReturnBook
	tempID := result[0].GetBorrowingID()
	borrowingRepo.ReturnBook(tempID)
	t.Logf("Calling ReturnBook()")

	// GetAllBorrowings
	result, _ = borrowingRepo.GetAllBorrowings()
	expected = make([]models.Borrowing, 1)
	t.Logf("Calling GetAllBorrowings(), result = %v\n", result)
	if len(result) != len(expected) {
		t.Errorf("ERROR. Expected: %v, got: %v", expected, result)
	}

	// SerializeBorrowingToXML
	XMLBorrowing, _ := xml_serializers.SerializeBorrowingToXML(result[0])
	t.Logf("Calling SerializeBorrowingToXML(), result = \n%v\n", string(XMLBorrowing))

	// DeserializeBorrowingFromXML
	borrowing, _ = xml_serializers.DeserializeBorrowingFromXML(XMLBorrowing)
	t.Logf("Calling DeserializeBorrowingFromXML(), result = %v\n", borrowing)

	// GetAllBorrowings
	result, _ = borrowingRepo.GetAllBorrowings()
	expected = make([]models.Borrowing, 1)
	t.Logf("Calling GetAllBorrowings(), result = %v\n", result)
	if len(result) != len(expected) {
		t.Errorf("ERROR. Expected: %v, got: %v", expected, result)
	}
}

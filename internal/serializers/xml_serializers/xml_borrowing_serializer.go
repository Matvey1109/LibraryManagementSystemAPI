package xml_serializers

import (
	"encoding/xml"

	"github.com/Matvey1109/LibraryManagementSystemCore/core/models"
)

type BorrowingXML struct {
	XMLName    xml.Name `xml:"borrowing"`
	ID         string   `xml:"id"`
	BookID     string   `xml:"bookId"`
	MemberID   string   `xml:"memberId"`
	BorrowYear int      `xml:"borrowYear"`
	ReturnYear int      `xml:"returnYear"`
}

func SerializeBorrowingToXML(book models.Borrowing) ([]byte, error) {
	borrowingXML := BorrowingXML{
		ID:         book.ID,
		BookID:     book.BookID,
		MemberID:   book.MemberID,
		BorrowYear: book.BorrowYear,
		ReturnYear: book.ReturnYear,
	}

	data, err := xml.MarshalIndent(borrowingXML, "", "  ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func DeserializeBorrowingFromXML(data []byte) (models.Borrowing, error) {
	var borrowingXML BorrowingXML
	err := xml.Unmarshal(data, &borrowingXML)
	if err != nil {
		return models.Borrowing{}, err
	}
	return models.Borrowing{
		ID:         borrowingXML.ID,
		BookID:     borrowingXML.BookID,
		MemberID:   borrowingXML.MemberID,
		BorrowYear: borrowingXML.BorrowYear,
		ReturnYear: borrowingXML.ReturnYear,
	}, nil
}

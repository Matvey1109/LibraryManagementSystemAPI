package xml_serializers

import (
	"encoding/xml"

	"github.com/Matvey1109/LibraryManagementSystemCore/core/models"
)

type BookXML struct {
	XMLName         xml.Name `xml:"book"`
	ID              string   `xml:"id"`
	Title           string   `xml:"title"`
	Author          string   `xml:"author"`
	PublicationYear int      `xml:"publicationYear"`
	Genre           string   `xml:"genre"`
	AvailableCopies int      `xml:"availableCopies"`
	TotalCopies     int      `xml:"totalCopies"`
}

func SerializeBookToXML(book models.Book) ([]byte, error) {
	bookXML := BookXML{
		ID:              book.ID,
		Title:           book.Title,
		Author:          book.Author,
		PublicationYear: book.PublicationYear,
		Genre:           book.Genre,
		AvailableCopies: book.AvailableCopies,
		TotalCopies:     book.TotalCopies,
	}

	data, err := xml.MarshalIndent(bookXML, "", "  ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func DeserializeBookFromXML(data []byte) (models.Book, error) {
	var bookXML BookXML
	err := xml.Unmarshal(data, &bookXML)
	if err != nil {
		return models.Book{}, err
	}
	return models.Book{
		ID:              bookXML.ID,
		Title:           bookXML.Title,
		Author:          bookXML.Author,
		PublicationYear: bookXML.PublicationYear,
		Genre:           bookXML.Genre,
		AvailableCopies: bookXML.AvailableCopies,
		TotalCopies:     bookXML.TotalCopies,
	}, nil
}

package json_serializers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Matvey1109/LibraryManagementSystemCore/core/models"
)

func SerializeBookToJson(book models.Book) ([]byte, error) {
	jsonData, err := json.Marshal(book)
	if err != nil {
		return nil, fmt.Errorf("error serializing book to JSON: %w", err)
	}
	formattedJson := bytes.Replace(jsonData, []byte(","), []byte(",\n"), -1)
	return formattedJson, nil
}

func DeserializeBookFromJson(jsonData []byte) (models.Book, error) {
	var book models.Book
	err := json.Unmarshal(jsonData, &book)
	if err != nil {
		return book, fmt.Errorf("error deserializing JSON to book: %w", err)
	}
	return book, nil
}

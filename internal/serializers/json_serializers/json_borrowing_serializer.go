package json_serializers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Matvey1109/LibraryManagementSystemAPI/internal/schemas"
)

func SerializeBorrowingToJson(borrowing schemas.Borrowing) ([]byte, error) {
	jsonData, err := json.Marshal(borrowing)
	if err != nil {
		return nil, fmt.Errorf("error serializing borrowing to JSON: %w", err)
	}
	formattedJson := bytes.Replace(jsonData, []byte(","), []byte(",\n"), -1)
	return formattedJson, nil
}

func DeserializeBorrowingFromJson(jsonData []byte) (schemas.Borrowing, error) {
	var borrowing schemas.Borrowing
	err := json.Unmarshal(jsonData, &borrowing)
	if err != nil {
		return borrowing, fmt.Errorf("error deserializing JSON to borrowing: %w", err)
	}
	return borrowing, nil
}

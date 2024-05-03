package serializers

import (
	"errors"
	"fmt"
	"math"
	"regexp"
)

func ValidateKeys(data map[string]interface{}, allowedKeys map[string]bool, requiredKeys bool) error {
	for key := range data {
		if _, ok := allowedKeys[key]; !ok {
			return fmt.Errorf("invalid key: %s", key)
		}
		allowedKeys[key] = true
	}
	if requiredKeys {
		for key, found := range allowedKeys {
			if !found {
				return fmt.Errorf("missing required key: %s", key)
			}
		}
	}
	return nil
}

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

func HasFraction(value float64) bool {
	return value != math.Trunc(value)
}

func ExtractString(data map[string]interface{}, key string) (string, error) {
	if value, ok := data[key].(string); ok {
		return value, nil
	}
	return "", errors.New("invalid input data")
}

func ExtractInt(data map[string]interface{}, key string) (int, error) {
	if value, ok := data[key].(float64); ok {
		if HasFraction(value) || value < 0 {
			return 0, errors.New("invalid input data")
		} else {
			return int(value), nil
		}
	}
	return 0, errors.New("invalid input data")
}

func ExtractOptionalString(data map[string]interface{}, key string) (*string, error) {
	var (
		res *string
		err error
	)
	if value, ok := data[key]; ok {
		if str, ok := value.(string); ok {
			res = &str
		} else if value == nil {
			res = nil
		} else {
			err = errors.New("invalid input data")
		}
	} else {
		res = nil
	}
	return res, err
}

func ExtractOptionalInt(data map[string]interface{}, key string) (*int, error) {
	var (
		res *int = new(int)
		err error
	)
	if value, ok := data[key]; ok {
		if num, ok := value.(float64); ok {
			if HasFraction(num) || num < 0 {
				err = errors.New("invalid input data")
			} else {
				*res = int(num)
			}
		} else if value == nil {
			res = nil
		} else {
			err = errors.New("invalid input data")
		}
	} else {
		res = nil
	}
	return res, err
}

func ValidateAddMemberData(data map[string]interface{}) (string, string, string, error) {
	allowedKeys := map[string]bool{
		"name":    false,
		"address": false,
		"email":   false,
	}

	if err := ValidateKeys(data, allowedKeys, true); err != nil {
		return "", "", "", err
	}

	var (
		name, address, email string
		err                  error
	)

	for key := range data {
		switch key {
		case "name":
			name, err = ExtractString(data, "name")
		case "address":
			address, err = ExtractString(data, "address")
		case "email":
			email, err = ExtractString(data, "email")
			if !ValidateEmail(email) {
				return "", "", "", errors.New("invalid email")
			}
		}
		if err != nil {
			return "", "", "", err
		}
	}

	return name, address, email, nil
}

func ValidateUpdateMemberData(data map[string]interface{}) (*string, *string, *string, error) {
	allowedKeys := map[string]bool{
		"name":    false,
		"address": false,
		"email":   false,
	}

	var (
		name, address, email *string
		err                  error
	)

	if err := ValidateKeys(data, allowedKeys, false); err != nil {
		return nil, nil, nil, err
	}

	for key := range data {
		switch key {
		case "name":
			name, err = ExtractOptionalString(data, "name")
		case "address":
			address, err = ExtractOptionalString(data, "address")
		case "email":
			email, err = ExtractOptionalString(data, "email")
			if !ValidateEmail(*email) {
				return nil, nil, nil, errors.New("invalid email")
			}
		}
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return name, address, email, nil
}

func ValidateAddBookData(data map[string]interface{}) (string, string, int, string, int, error) {
	allowedKeys := map[string]bool{
		"title":           false,
		"author":          false,
		"publicationYear": false,
		"genre":           false,
		"totalCopies":     false,
	}

	if err := ValidateKeys(data, allowedKeys, true); err != nil {
		return "", "", 0, "", 0, err
	}

	var (
		title, author, genre         string
		publicationYear, totalCopies int
		err                          error
	)

	for key := range data {
		switch key {
		case "title":
			title, err = ExtractString(data, "title")
		case "author":
			author, err = ExtractString(data, "author")
		case "publicationYear":
			publicationYear, err = ExtractInt(data, "publicationYear")
		case "genre":
			genre, err = ExtractString(data, "genre")
		case "totalCopies":
			totalCopies, err = ExtractInt(data, "totalCopies")
		}
		if err != nil {
			return "", "", 0, "", 0, err
		}
	}

	return title, author, publicationYear, genre, totalCopies, nil
}

func ValidateUpdateBookData(data map[string]interface{}) (*string, *string, *int, *string, *int, *int, error) {
	allowedKeys := map[string]bool{
		"title":           false,
		"author":          false,
		"publicationYear": false,
		"genre":           false,
		"availableCopies": false,
		"totalCopies":     false,
	}

	var (
		title, author, genre                          *string
		publicationYear, availableCopies, totalCopies *int
		err                                           error
	)

	if err := ValidateKeys(data, allowedKeys, false); err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	for key := range data {
		switch key {
		case "title":
			title, err = ExtractOptionalString(data, "title")
		case "author":
			author, err = ExtractOptionalString(data, "author")
		case "publicationYear":
			publicationYear, err = ExtractOptionalInt(data, "publicationYear")
		case "genre":
			genre, err = ExtractOptionalString(data, "genre")
		case "availableCopies":
			availableCopies, err = ExtractOptionalInt(data, "availableCopies")
		case "totalCopies":
			totalCopies, err = ExtractOptionalInt(data, "totalCopies")
		}
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
	}

	return title, author, publicationYear, genre, availableCopies, totalCopies, nil
}

func ValidateAddBorrowingData(data map[string]interface{}) (string, string, int, error) {
	allowedKeys := map[string]bool{
		"memberID":   false,
		"bookID":     false,
		"borrowYear": false,
	}

	if err := ValidateKeys(data, allowedKeys, true); err != nil {
		return "", "", 0, err
	}

	var (
		bookID, memberID string
		borrowYear       int
		err              error
	)

	for key := range data {
		switch key {
		case "bookID":
			bookID, err = ExtractString(data, "bookID")
		case "memberID":
			memberID, err = ExtractString(data, "memberID")
		case "borrowYear":
			borrowYear, err = ExtractInt(data, "borrowYear")
		}
		if err != nil {
			return "", "", 0, err
		}
	}

	return bookID, memberID, borrowYear, nil
}

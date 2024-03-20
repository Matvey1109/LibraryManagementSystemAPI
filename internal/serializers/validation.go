package serializers

import (
	"errors"
	"fmt"
	"math"
	"regexp"
)

func validateKeys(data map[string]interface{}, allowedKeys map[string]bool, requiredKeys bool) error {
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

func validateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

func hasFraction(value float64) bool {
	return value != math.Trunc(value)
}

func extractString(data map[string]interface{}, key string) (string, error) {
	if value, ok := data[key].(string); ok {
		return value, nil
	}
	return "", errors.New("invalid input data")
}

func extractInt(data map[string]interface{}, key string) (int, error) {
	if value, ok := data[key].(float64); ok {
		if hasFraction(value) || value < 0 {
			return 0, errors.New("invalid input data")
		} else {
			return int(value), nil
		}
	}
	return 0, errors.New("invalid input data")
}

func extractOptionalString(data map[string]interface{}, key string) (*string, error) {
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

func extractOptionalInt(data map[string]interface{}, key string) (*int, error) {
	var (
		res *int = new(int)
		err error
	)
	if value, ok := data[key]; ok {
		if num, ok := value.(float64); ok {
			if hasFraction(num) || num < 0 {
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

	if err := validateKeys(data, allowedKeys, true); err != nil {
		return "", "", "", err
	}

	var (
		name, address, email string
		err                  error
	)

	for key := range data {
		switch key {
		case "name":
			name, err = extractString(data, "name")
		case "address":
			address, err = extractString(data, "address")
		case "email":
			email, err = extractString(data, "email")
			if !validateEmail(email) {
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

	if err := validateKeys(data, allowedKeys, false); err != nil {
		return nil, nil, nil, err
	}

	for key := range data {
		switch key {
		case "name":
			name, err = extractOptionalString(data, "name")
		case "address":
			address, err = extractOptionalString(data, "address")
		case "email":
			email, err = extractOptionalString(data, "email")
			if !validateEmail(*email) {
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

	if err := validateKeys(data, allowedKeys, true); err != nil {
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
			title, err = extractString(data, "title")
		case "author":
			author, err = extractString(data, "author")
		case "publicationYear":
			publicationYear, err = extractInt(data, "publicationYear")
		case "genre":
			genre, err = extractString(data, "genre")
		case "totalCopies":
			totalCopies, err = extractInt(data, "totalCopies")
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

	if err := validateKeys(data, allowedKeys, false); err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	for key := range data {
		switch key {
		case "title":
			title, err = extractOptionalString(data, "title")
		case "author":
			author, err = extractOptionalString(data, "author")
		case "publicationYear":
			publicationYear, err = extractOptionalInt(data, "publicationYear")
		case "genre":
			genre, err = extractOptionalString(data, "genre")
		case "availableCopies":
			availableCopies, err = extractOptionalInt(data, "availableCopies")
		case "totalCopies":
			totalCopies, err = extractOptionalInt(data, "totalCopies")
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

	if err := validateKeys(data, allowedKeys, true); err != nil {
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
			bookID, err = extractString(data, "bookID")
		case "memberID":
			memberID, err = extractString(data, "memberID")
		case "borrowYear":
			borrowYear, err = extractInt(data, "borrowYear")
		}
		if err != nil {
			return "", "", 0, err
		}
	}

	return bookID, memberID, borrowYear, nil
}

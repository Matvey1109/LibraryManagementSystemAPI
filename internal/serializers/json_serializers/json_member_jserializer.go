package json_serializers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Matvey1109/LibraryManagementSystemCore/core/models"
)

func SerializeMemberToJson(member models.Member) ([]byte, error) {
	jsonData, err := json.Marshal(member)
	if err != nil {
		return nil, fmt.Errorf("error serializing member to JSON: %w", err)
	}
	formattedJson := bytes.Replace(jsonData, []byte(","), []byte(",\n"), -1)
	return formattedJson, nil
}

func DeserializeMemberFromJson(jsonData []byte) (models.Member, error) {
	var member models.Member
	err := json.Unmarshal(jsonData, &member)
	if err != nil {
		return member, fmt.Errorf("error deserializing JSON to member: %w", err)
	}
	return member, nil
}

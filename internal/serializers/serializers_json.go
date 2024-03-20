package serializers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SerializeJsonData(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error serializing data to JSON: %w", err)
	}
	formattedJson := bytes.Replace(jsonData, []byte(","), []byte(",\n"), -1)
	return formattedJson, nil
}

func DeserializeJsonData(r *http.Request) (map[string]interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{}
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("error deserializing JSON data: %w", err)
	}
	return data, nil
}

package utils

import "encoding/json"

func EncodeToJsonString(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func DecodeFromJsonString(bytes []byte, obj interface{}) error {
	return json.Unmarshal(bytes, obj)
}

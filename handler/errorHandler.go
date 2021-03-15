package handler

import (
	"encoding/json"
	"errors"
	"fmt"
)

func GetResponseError(data []byte) error {
	errMessage := make(map[string]interface{})
	json.Unmarshal([]byte(string(data)), &errMessage)

	return errors.New(fmt.Sprintf("%v", errMessage["message"]))
}

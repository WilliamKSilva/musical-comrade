package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RequestJSON(req *http.Request, genericStruct interface{}) error {
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(genericStruct)

	return err
}

func ResponseJSON(w io.Writer, genericStruct interface{}) error {
	encoder := json.NewEncoder(w)

	err := encoder.Encode(genericStruct)

	return err
}

func FieldValidator(w http.ResponseWriter, genericStruct interface{}, fields []string) (string, error) {
	var mymap map[string]interface{}
	data, err := json.Marshal(genericStruct)

	json.Unmarshal(data, &mymap)

	for i := 0; i < len(fields); i++ {
		valid := mymap[fields[i]]

		if valid == "" {
			message := fmt.Sprintf("The field: %s, cannot be empty", fields[i])

			return message, nil
		}
	}

	return "", err
}

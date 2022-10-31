package presentation

import (
	"encoding/json"
	"io"
	"net/http"
)

func requestJSON(req *http.Request, genericStruct interface{}) error {
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(genericStruct)

	return err
}

func responseJSON(w io.Writer, genericStruct interface{}) error {
	encoder := json.NewEncoder(w)

	err := encoder.Encode(genericStruct)

	return err
}

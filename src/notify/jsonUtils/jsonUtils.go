package jsonUtils

import (
	"encoding/json"
	"io"
	"net/http"
)

type jsonError struct {
	HTTPStatus int    `json:"http-status"`
	Message    string `json:"message"`
}

// Output takes data, marshals it into JSON and writes it out to the http.ResponseWriter
func Output(w http.ResponseWriter, httpStatus int, data interface{}) {
	eJSON, err := json.Marshal(data)
	if err != nil {
		write500Error(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(eJSON)
}

// Decode will try and parse the content of a Reader into the matching interface
func Decode(body io.Reader, s interface{}) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&s)
	return err
}

// Error takes an error string and writes it out to the http.ResponseWriter
func Error(w http.ResponseWriter, httpStatus int, errorStr string) {
	data := jsonError{httpStatus, errorStr}

	Output(w, httpStatus, data)
}

func write500Error(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{\"http-status\": 500, \"message\":\"Internal Error\"}"))
}

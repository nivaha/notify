package jsonUtils

import (
	"encoding/json"
	"net/http"
)

type jsonError struct {
	HTTPStatus int    `json:"http-status"`
	Message    string `json:"message"`
}

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

func Error(w http.ResponseWriter, httpStatus int, errorStr string) {
	data := jsonError{httpStatus, errorStr}

	Output(w, httpStatus, data)
}

func write500Error(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{\"http-status\": 500, \"message\":\"Internal Error\"}"))
}

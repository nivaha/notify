package jsonUtils

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonError struct {
	HTTPStatus int    `json:"http-status"`
	Message    string `json:"message"`
}

func Output(w http.ResponseWriter, httpStatus int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	eJSON, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"http-status\": 500, \"message\":\"Internal Error\"}"))
		return
	}

	w.WriteHeader(httpStatus)
	w.Write(eJSON)
}

func Error(w http.ResponseWriter, httpStatus int, errorStr string) {
	w.Header().Set("Content-Type", "application/json")

	data := jsonError{httpStatus, errorStr}

	mapB, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"http-status\": 500, \"message\":\"Internal Error\"}"))
		return
	}

	log.Printf(string(mapB))

	w.WriteHeader(httpStatus)
	w.Write(mapB)
}

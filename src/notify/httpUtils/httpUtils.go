package httpUtils

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonError struct {
	HTTPStatus int    `json:"http-status"`
	Message    string `json:"message"`
}

func ErrorJSON(w http.ResponseWriter, httpStatus int, errorStr string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	data := jsonError{httpStatus, errorStr}

	mapB, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte("{\"http-status\": 500, \"message\":\"Internal Error\"}"))
		return
	}

	log.Printf(string(mapB))

	w.Write(mapB)
}

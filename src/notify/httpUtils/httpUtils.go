package httpUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type jsonError struct {
	HTTPStatus int    `json:"http-status"`
	Message    string `json:"message"`
}

func ErrorJSON(w http.ResponseWriter, httpStatus int, errorStr string) {
	data := jsonError{httpStatus, errorStr}

	mapB, _ := json.Marshal(data)
	log.Printf(string(mapB))

	w.WriteHeader(httpStatus)
	fmt.Fprintf(w, string(mapB))
}

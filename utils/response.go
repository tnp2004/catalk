package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type response struct {
	Data any `json:"data"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func SuccessResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := &response{Data: data}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "	")
	if err := encoder.Encode(resp); err != nil {
		log.Printf("error marshal response body. Err: %s", err.Error())
		ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("marshal body error"))
	}
}

func ErrorResponse(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := &errorResponse{Message: err.Error()}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "	")
	if err := encoder.Encode(resp); err != nil {
		log.Printf("error marshal error response body. Err: %s", err.Error())
		ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("marshal body error"))
	}
}

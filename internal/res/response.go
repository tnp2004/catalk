package res

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
	w.WriteHeader(status)
	resp := &response{Data: data}
	jsonResp, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		log.Printf("error marshal response body. Err: %s", err.Error())
		ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("marshal body error"))
	}

	w.Write(jsonResp)
}

func ErrorResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	errResp := &errorResponse{Message: err.Error()}
	jsonResp, err := json.MarshalIndent(errResp, "", "\t")
	if err != nil {
		log.Printf("error marshal error response body. Err: %s", err.Error())
		ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("marshal body error"))
	}

	w.Write(jsonResp)
}

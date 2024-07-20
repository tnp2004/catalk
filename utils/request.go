package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ReadReqBody(r *http.Request, data any) error {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error read body. Err: %s", err.Error())
		return fmt.Errorf("read body failed")
	}

	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("error handling JSON marshal. Err: %s", err.Error())
		return fmt.Errorf("handling JSON marshal failed")
	}

	return nil
}

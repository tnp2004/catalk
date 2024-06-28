package server

import (
	"catalk/internal/ai"
	"catalk/internal/res"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (s *Server) ChatWithGeminiHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	reqBody := new(ai.GeminiTextBody)
	if err := json.Unmarshal(body, &reqBody); err != nil {
		log.Printf("error unmarshal body. Err: %s", err.Error())
		res.ErrorResponse(w, http.StatusInternalServerError, err)
	}

	resp, err := ai.TextToGemini(reqBody)
	if err != nil {
		res.ErrorResponse(w, http.StatusInternalServerError, err)
	}

	res.SuccessResponse(w, http.StatusOK, resp)
}

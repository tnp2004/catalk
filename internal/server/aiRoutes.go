package server

import (
	"catalk/internal/ai"
	"catalk/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (s *Server) ChatWithGeminiHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	reqBody := new(ai.GeminiRequest)
	if err := json.Unmarshal(body, &reqBody); err != nil {
		log.Printf("error unmarshal body. Err: %s", err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
	}

	resp, err := ai.TextToGemini(reqBody)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
	}

	utils.SuccessResponse(w, http.StatusOK, resp)
}

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
	breed := r.PathValue("breed")

	if err := json.Unmarshal(body, &reqBody); err != nil {
		log.Printf("error handling JSON marshal. Err: %s", err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp, err := ai.TextToGemini(reqBody, breed)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(w, http.StatusOK, resp)
}

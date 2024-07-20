package server

import (
	"catalk/internal/ai"
	"catalk/utils"
	"net/http"
)

func (s *Server) ChatWithGeminiHandler(w http.ResponseWriter, r *http.Request) {
	reqBody := new(ai.GeminiRequest)
	if err := utils.ReadReqBody(r, reqBody); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	breed := r.PathValue("breed")
	ai := ai.NewAi(s.config.Google)
	resp, err := ai.TextToGemini(reqBody, breed)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(w, http.StatusOK, resp)
}

package server

import (
	"catalk/internal/ai"
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
	}

	resp, err := ai.TextToGemini(reqBody)
	if err != nil {
		log.Println(err.Error())
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Println(err.Error())
	}

	w.Write(jsonResp)
}

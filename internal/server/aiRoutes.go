package server

import (
	"catalk/internal/ai"
	"encoding/json"
	"fmt"
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
	fmt.Printf("%#v", resp)
}

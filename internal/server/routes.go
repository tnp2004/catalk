package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/", s.apiV1())

	return mux
}

func (s *Server) apiV1() http.Handler {
	v1 := http.NewServeMux()
	v1.Handle("/", http.HandlerFunc(s.HelloWorldHandler))
	v1.Handle("/health", http.HandlerFunc(s.healthHandler))
	v1.Handle("POST /gemini/cats/{breed}", http.HandlerFunc(s.ChatWithGeminiHandler))

	return http.StripPrefix("/api/v1", v1)
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

package server

import (
	"catalk/utils"
	"encoding/json"
	"fmt"
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
	v1.Handle("/", http.HandlerFunc(s.NotMatchRoutes))
	v1.Handle("GET /db/health", http.HandlerFunc(s.dbHealthHandler))
	v1.Handle("GET /server/health", http.HandlerFunc(s.ServerHealthHandler))

	v1.Handle("GET /cats/breeds", http.HandlerFunc(s.CatBreeds))
	v1.Handle("POST /gemini/cats/{breed}", http.HandlerFunc(s.ChatWithGeminiHandler))

	//auth
	v1.Handle("POST /auth/google/login", http.HandlerFunc(s.GoogleLogin))
	v1.Handle("POST /auth/google/callback", http.HandlerFunc(s.GoogleCallback))

	return http.StripPrefix("/api/v1", v1)
}

func (s *Server) NotMatchRoutes(w http.ResponseWriter, r *http.Request) {
	utils.ErrorResponse(w, http.StatusNotFound, fmt.Errorf("route not found"))
}

func (s *Server) ServerHealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "OK"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("error handling JSON marshal. Err: %s", err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (s *Server) dbHealthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Printf("error handling JSON marshal. Err: %s", err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

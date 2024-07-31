package server

import (
	"catalk/internal/auth/google"
	"catalk/utils"
	"fmt"
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

	// auth
	googleOAuth := google.NewGoogleOAuth(google.GoogleConfig(), s.config.Google, s.config.Database, s.config.JWT)
	v1.Handle("GET /auth/google/login", http.HandlerFunc(googleOAuth.GoogleLoginHandler))
	v1.Handle("GET /auth/google/callback", http.HandlerFunc(googleOAuth.GoogleCallbackHandler))

	return http.StripPrefix("/api/v1", v1)
}

func (s *Server) NotMatchRoutes(w http.ResponseWriter, r *http.Request) {
	utils.ErrorResponse(w, http.StatusNotFound, fmt.Errorf("route not found"))
}

func (s *Server) ServerHealthHandler(w http.ResponseWriter, r *http.Request) {
	utils.MessageResponse(w, http.StatusOK, "server health ok")
}

func (s *Server) dbHealthHandler(w http.ResponseWriter, r *http.Request) {
	utils.SuccessResponse(w, http.StatusOK, s.db.Health())
}

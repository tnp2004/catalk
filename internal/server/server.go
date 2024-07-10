package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"

	"catalk/internal/database"
)

type Server struct {
	port int
	db   database.Service
}

type ServerWrapper struct {
	hostName   string
	httpServer *http.Server
}

func NewServer() *ServerWrapper {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db: database.New(),
	}

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4321"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	})

	// Declare Server config
	server := &ServerWrapper{
		hostName: os.Getenv("HOST_NAME"),
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", NewServer.port),
			Handler:      cors.Handler(NewServer.RegisterRoutes()),
			IdleTimeout:  time.Minute,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	}

	return server
}

func (s *ServerWrapper) Start() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefulShutdown(ch)

	s.httpListening()
}

func (s *ServerWrapper) httpListening() {
	log.Printf("Server listening on %s%s%s", "http://", s.hostName, s.httpServer.Addr)
	err := s.httpServer.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Printf("error listening server. Err: %s", err.Error())
	}
}

func (s *ServerWrapper) gracefulShutdown(ch chan os.Signal) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	<-ch
	log.Println("Server gracefully shutdown")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("error while shutting down server. Err: %s", err.Error())
	}
}

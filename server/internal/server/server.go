package server

import (
	"fmt"
	"log"
	"net/http"
	"renting/internal/config"
	"renting/internal/handlers"
)

type Server struct {
	c    *config.Config
	h    *handlers.Handlers
	port string
}

func NewServer(config *config.Config, port string) *Server {
	return &Server{
		c:    config,
		port: port,
	}
}

func (s *Server) Run() {
	s.h = handlers.NewHandlers(s.c)
	router := s.router()

	fmt.Println("listening port", s.port)
	err := http.ListenAndServe(s.port, router)
	if err != nil {
		log.Fatal("can't listen server")
	}
}

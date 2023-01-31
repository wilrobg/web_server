package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	return http.ListenAndServe(s.port, nil)
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}

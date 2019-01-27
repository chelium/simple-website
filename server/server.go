package server

import (
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"

	"github.com/chelium/simple-website/todo"
)

// Server holds the dependencies for an HTTP server.
type Server struct {
	Todo todo.todoService

	router *gin.Engine
}

// New returns a new HTTP server.
func New(ts todo.todoService) *Server {
	s := &Server{
		Todo: ts,
	}

	r := gin.Default()

	todoHandler := &todoHandler{
		s: ts,
	}

	todoHandler.AddRoutes(r)
	s.router = r
	return s
}

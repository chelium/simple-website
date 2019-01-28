package server

import (
	"github.com/gin-gonic/gin"
	// jose "gopkg.in/square/go-jose.v2"

	todo "github.com/chelium/simple-website/todo/service"
	user "github.com/chelium/simple-website/user/service"
)

// Server holds the dependencies for an HTTP server.
type Server struct {
	Todo todo.Service
	User user.Service

	router *gin.Engine
}

// New returns a new HTTP server.
func New(ts todo.Service, us user.Service) *Server {
	s := &Server{
		Todo: ts,
		User: us,
	}

	r := gin.Default()

	todoHandler := &todoHandler{
		s: ts,
	}
	todoHandler.AddRoutes(r)

	userHandler := &userHandler{
		s: us,
	}
	userHandler.AddRoutes(r)

	s.router = r
	return s
}

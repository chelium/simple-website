package server

import (
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	s := &Server{}
	r := gin.Default()
	s.router = r
	return s
}

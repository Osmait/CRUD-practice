package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr    string
	engine      *gin.Engine
	PostService PostService
}

func NewServer(host string, port uint, postService PostService) Server {
	srv := Server{
		engine:      gin.New(),
		httpAddr:    fmt.Sprintf("%s:%d", host, port),
		PostService: postService,
	}
	srv.routes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) routes() {
	s.engine.GET("/", FindPost(s.PostService))
	s.engine.POST("/", CreatedPost(s.PostService))
}

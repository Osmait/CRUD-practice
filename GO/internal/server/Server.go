package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

type Server struct {
	httpAddr        string
	engine          *gin.Engine
	PostService     PostService
	shutdownTimeout time.Duration
}

func NewServer(ctx context.Context, host string, port uint, postService PostService, shutdownTimeout time.Duration) (context.Context, Server) {
	srv := Server{
		engine:          gin.Default(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		PostService:     postService,
		shutdownTimeout: shutdownTimeout,
	}
	srv.routes()
	return serverContext(ctx), srv
}

func (s *Server) routes() {

	s.engine.GET("/", FindPost(s.PostService))
	s.engine.POST("/", CreatedPost(s.PostService))

}
func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)
	handle := cors.AllowAll().Handler(s.engine)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: handle,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}

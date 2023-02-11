package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var ProviderSet = wire.NewSet(
	New,
	NewRouter,
	wire.Bind(new(http.Handler), new(*httprouter.Router)),
)

type Server struct {
	httpServer   *http.Server
	Port         uint
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *Server) Name() string {
	return "http.Server"
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func New(o *Option, handler http.Handler) *Server {
	addr := fmt.Sprintf(":%d", o.Port)
	server := &http.Server{
		Handler:      handler,
		Addr:         addr,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
	}
	return &Server{
		httpServer:   server,
		Port:         o.Port,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
	}
}

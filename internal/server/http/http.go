package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func New(o *Option) (*http.ServeMux, func()) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello world"))
	})
	server := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%d", o.Port),
	}
	go func() {
		log.Println("server is starting...")
		server.ListenAndServe()
	}()
	cleanup := func() {
		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
		server.Shutdown(ctx)
		log.Println("server is closed")
	}
	return mux, cleanup
}

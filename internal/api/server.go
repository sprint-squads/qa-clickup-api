package api

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Sprint-Squads/qa-clickup-api/pkg/application"
	"net/http"
)

// Server strcut
type Server struct {
	*http.Server
}

// NewServer - creates Server instance
func NewServer() (*Server, error) {
	fmt.Println("Configuring server..")

	app, err := application.Get()
	if err != nil {
		return nil, err
	}

	fmt.Println("host", app.Config.Server.Host, "port", app.Config.Server.Port)
	router, err := New(*app)
	if err != nil {
		return nil, err
	}
	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", app.Config.Server.Host, app.Config.Server.Port),
		Handler: router,
	}

	return &Server{&srv}, nil
}

// Start - starts server
func (srv *Server) Start() {
	fmt.Println("Starting server..")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	fmt.Println("Listening on", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	fmt.Println("Shutting down server.. Reason:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("Server gracefully stopped")
}

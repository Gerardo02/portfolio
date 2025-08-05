package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gerardo02/porfolio/internal/handler"
	"github.com/gerardo02/porfolio/internal/handler/home"
	"github.com/gin-gonic/gin"
)

func Start() error {
	r := gin.Default()

	r.Static("/static", "./public")

	handler := handler.Handler{
		Home: home.HomeHandler{},
	}

	NewRouter(r, handler)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	errChan := make(chan error, 1)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Ready and listening for connections...")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	select {
	case sig := <-stopChan:
		log.Printf("Recieved signal: %v", sig)
	case err := <-errChan:
		return err
	}

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error when shutting down the server: %v", err)
		return err
	}

	log.Println("Server exited cleanly.")

	return nil
}

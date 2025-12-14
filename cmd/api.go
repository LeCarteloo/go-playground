package main

import (
	"go_playground/internal/posts"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) mount() http.Handler {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.RequestID) // useful for rate limiting
	router.Use(middleware.RealIP)    // useful for rate limiting and analytics
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer) // recover from crashes (aka panics)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second)) // 60sec limit for each request

	router.Get("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	})

	postsHandler := posts.NewHandler(nil)

	router.Get("/posts", postsHandler.ListPosts)

	return router
}

func (app *application) run(handler http.Handler) error {
	server := http.Server{
		Addr:         app.config.addr,
		Handler:      handler,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Starting server on %v", app.config.addr)

	return server.ListenAndServe()

}

type application struct {
	config config
	// logger
	// db driver
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{}))
}

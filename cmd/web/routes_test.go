package main

import (
	"github.com/fzambone/go-bookings-app/internal/config"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing; test passed
	default:
		t.Errorf("Type is not *chi.mux, type is %T", v)
	}
}

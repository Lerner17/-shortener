package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Lerner17/shortener/internal/db"
	mw "github.com/Lerner17/shortener/internal/middleware"
)

func NewRouter() chi.Router {

	var db = db.GetDB()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mw.GZIPMiddleware)
	r.Post("/api/shorten", ShortenerAPIHandler(db))
	r.Get("/{urlID}", RedirectHandler(db))
	r.Post("/", CreateShortURLHandler(db))
	return r
}

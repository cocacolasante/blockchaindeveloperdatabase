package application

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *Application) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	// mux.Use(app.EnableCORS)

	mux.Post("/signup", app.CreateWalletAccount)

	return mux
}

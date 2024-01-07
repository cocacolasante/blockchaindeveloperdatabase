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

	mux.Get("/wallet/{address}", app.GetWalletAccount)

	// how to protect/auth signup so randoms cant create accounts for other users wallet addresses
	mux.Post("/signup", app.CreateWalletAccount)
	// mux.Post("/login", )

	// mux.Post("/refreshapikey", )

	return mux
}

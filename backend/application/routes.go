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

	// protected routes -uses api key
	mux.Route("/{address}", func(muxx chi.Router) {
		muxx.Use(app.authRequired)
		muxx.Get("/refreshapikey", app.RefreshApiKey)
		muxx.Get("/", app.GetAccountFromDatabase)
		// add smart contract to a users account   ---- needs to add in the smart contract call to deduct credits
		muxx.Post("/newcontract", app.AddSmartContractToAccount)
	})
	// mux.Post("/login", )
	// create a handler to get api key from login


	// create admin routes for myself


	return mux
}



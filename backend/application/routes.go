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
	mux.Use(app.EnableCORS)

	mux.Get("/wallet/{address}", app.GetWalletAccount)

	// how to protect/auth signup so randoms cant create accounts for other users wallet addresses
	mux.Post("/signup", app.CreateWalletAccount)

	// create a handler to get api key from login credentials
	mux.Post("/login", app.LoginWithEmail) // use to get api key with login credentials

	// get remaining credits
	
	mux.Post("/apikeywithemail", app.APIKeyWithLogin)

	// validate apikey against database for frontend
	mux.Post("/validatekey", app.ValidateApiKey)


	// protected routes -uses api key
	mux.Route("/{address}", func(muxx chi.Router) {
		muxx.Use(app.authRequired)
		muxx.Use(app.ActiveAccountMiddleware)
		muxx.Get("/refreshapikey", app.RefreshApiKey)
		muxx.Get("/", app.GetAccountFromDatabase)
		// add smart contract to a users account   ---- needs to add in the smart contract call to deduct credits
		muxx.Post("/newcontract", app.AddSmartContractToAccount)
		muxx.Get("/contract/{contractaddress}", app.GetSmartContract)
		muxx.Patch("/contract/{contractaddress}", app.UpdateSmartContract)
		muxx.Delete("/contract/{contractaddress}", app.DeleteSmartContract)

		muxx.Get("/contracts", app.GetAllSmartContractAddressesByWallet)
		muxx.Get("/fullcontracts", app.GetSmartContractFullByWallet)


		muxx.Get("/credits", app.GetRemainCreditsByAddress)
		
		
	})

	mux.Route("/activate/{address}", func(r chi.Router) {
		r.Use(app.EmailAuthMiddleware)
		r.Get("/", app.ActivateAccount)

	})


	mux.Route("/admin", func(r chi.Router) {
		
	})

	// create admin routes for myself


	return mux
}



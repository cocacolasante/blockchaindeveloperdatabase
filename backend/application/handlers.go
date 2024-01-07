package application

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/models"
	"github.com/go-chi/chi/v5"
)

func (app *Application) CreateWalletAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	// add in a password hash field

	var input models.WalletAccount
	err := app.ReadJSON(w, r, &input)

	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}

	existing, err := app.DB.GetWalletByAddress(input.WalletAddress)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}

	if existing.WalletAddress != ""{
		app.ErrorJSON(w, errors.New("account already created in database"))
		return
	}
	isValidated := app.ValidateSignUp(&input)
	if !isValidated {
		app.ErrorJSON(w, errors.New("missing fields in sign up creation"))
		return
	}
	hasPass := app.HashPassword(input.Password)
	log.Println(hasPass)
	// add wallet to db returns the new users api key to access api 
	apikey, err := app.DB.AddWalletToDb(input.WalletAddress, input.Email, hasPass)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}
	input.ApiKey = apikey
	input.Password = hasPass
	
	err = app.writeJSON(w, http.StatusAccepted, input)
	if err != nil {
		// Handle error writing JSON
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *Application) GetWalletAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.ErrorJSON(w, errors.New("invalid method"), http.StatusBadRequest)
		return
	}

	address := chi.URLParam(r, "address")
	app.InfoLog.Println(address)

	wallet, err := app.DB.GetWalletByAddress(address)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	currentCredits, err := app.Web3.GetRemainingCredits(wallet.WalletAddress)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	wallet.CreditsAvailable = currentCredits

	out, err := json.Marshal(wallet)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)

}

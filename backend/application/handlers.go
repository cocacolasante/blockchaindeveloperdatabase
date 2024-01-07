package application

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/models"
	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/tools"
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

	if existing.WalletAddress != "" {
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

// AUTHENTICATED HANDLERS

func (app *Application) RefreshApiKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.ErrorJSON(w, errors.ErrUnsupported, http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "address")
	newApiKey := tools.GenerateApiKey()
	newKey, err := app.DB.UpdateAPIKey(id, newApiKey)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadGateway)
		return
	}

	wallet := models.WalletAccount{
		WalletAddress: id,
		ApiKey:        newKey,
	}

	out, err := json.Marshal(wallet)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(out)

}

// GET USERS ACCOUNT FROM DATABASE
func (app *Application) GetAccountFromDatabase(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.ErrorJSON(w, errors.ErrUnsupported, http.StatusBadRequest)
		return
	}
	id := chi.URLParam(r, "address")

	wallet, err := app.DB.AdminGetWalletAccount(id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	out, err := json.Marshal(wallet)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)

}

func (app *Application) AddSmartContractToAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.ErrorJSON(w, errors.ErrUnsupported, http.StatusBadRequest)
		return
	}

	var smartContract models.SmartContract
	err := app.ReadJSON(w, r, &smartContract)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	if smartContract.DeployerWallet == "" {
		id := chi.URLParam(r, "address")
		smartContract.DeployerWallet = id
	}
	app.InfoLog.Println("smart contract from request", smartContract)

	err = app.DB.AddSmartContractToAccountDb(smartContract)

	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	
	var jsonPayload = struct {
		Message string `json:"message"`
		SmartContract models.SmartContract `json:"smart_contract"`
	}{
		Message: "Successfully Added To Database",
		SmartContract: smartContract,
	}
	out, err := json.Marshal(jsonPayload)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	w.Write(out)
}


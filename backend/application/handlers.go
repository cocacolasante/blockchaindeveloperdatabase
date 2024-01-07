package application

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/models"
	"github.com/go-chi/chi/v5"
)

func (app *Application) CreateWalletAccount(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	var input models.WalletAccount
	err := app.ReadJSON(w, r, &input)

	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}

	apikey, err := app.DB.AddWalletToDb(input.WalletAddress)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}
	input.ApiKey = apikey
	app.InfoLog.Println("hit in create wallet json")
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
	app.InfoLog.Println(wallet)

	out, err := json.Marshal(wallet)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)

}

package application

import (
	"net/http"

	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/models"
)

func (app *Application) CreateWalletAccount(w http.ResponseWriter, r *http.Request) {
	app.InfoLog.Println("hit")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}
	app.InfoLog.Println("hit")
	var input models.WalletAccount
	err := app.ReadJSON(w, r, &input)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}

	wall, err := app.DB.AddWalletToDb(input.WalletAddress)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusAccepted, wall)
	if err != nil {
		// Handle error writing JSON
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

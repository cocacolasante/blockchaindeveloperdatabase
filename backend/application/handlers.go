package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/models"
	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/tools"
	"github.com/go-chi/chi/v5"
)
type ApiKeyRequest struct {
    Email  string `json:"email"`
    ApiKey string `json:"api_key"`
}
// CREATE A NEW WALLET ACCOUNT IN DATABASE
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

	input.Active = false

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
	input.Password = hasPass
	// add wallet to db returns the new users api key to access api
	apikey, err := app.DB.AddWalletToDb(&input)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}

	input.ApiKey = apikey

	input.CreditsAvailable = big.NewInt(0)

	datamap := make(map[string]any)
	datamap["activatelink"] = fmt.Sprintf("http://%s:%d/activate/%s?key=%s", app.Domain, app.Port, input.WalletAddress, apikey)
	datamap["apikey"] = apikey

	msg, err := app.Mailer.CreateMessage(input.Email, "Please Activate Your Account", "confirmation-email", datamap)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}

	err = app.Mailer.SendEmail(msg)
	if err != nil {
		app.ErrorLog.Println(err)
		// Handle error writing JSON
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = app.writeJSON(w, http.StatusAccepted, input)
	if err != nil {
		// Handle error writing JSON
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// GET USER WALLET BY ADDRESS FROM DATABASE
func (app *Application) GetWalletAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.ErrorJSON(w, errors.New("invalid method"), http.StatusBadRequest)
		return
	}

	address := chi.URLParam(r, "address")

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

// ADD SMART CONTRACT ADDRESS TO WALLET ACCOUNT
func (app *Application) AddSmartContractToAccount(w http.ResponseWriter, r *http.Request) {
	app.InfoLog.Println("hit")
	if r.Method != http.MethodPost {
		app.ErrorJSON(w, errors.ErrUnsupported, http.StatusBadRequest)
		return
	}
	id := chi.URLParam(r, "address")

	balance, err := app.Web3.GetRemainingCredits(id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	if balance == big.NewInt(0) {
		app.ErrorJSON(w, errors.New("insufficient balance"))
		return
	}

	var smartContract models.SmartContract
	err = app.ReadJSON(w, r, &smartContract)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	if smartContract.DeployerWallet == "" {
		smartContract.DeployerWallet = id
	}

	app.InfoLog.Println(smartContract.Address)
	app.InfoLog.Println(smartContract.DeployerWallet)
	app.InfoLog.Println(smartContract.ProjectName)
	app.InfoLog.Println(smartContract.Description)

	err = app.Web3.RedeemCredits(id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}



	err = app.DB.AddSmartContractToAccountDb(smartContract, id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	var jsonPayload = struct {
		Message       string               `json:"message"`
		SmartContract models.SmartContract `json:"smart_contract"`
	}{
		Message:       "Successfully Added To Database",
		SmartContract: smartContract,
	}
	out, err := json.Marshal(jsonPayload)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	w.Write(out)
}

// GET SMART CONTRACT FROM DATABASE BY ADDRESS
func (app *Application) GetSmartContract(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.ErrorJSON(w, errors.ErrUnsupported, http.StatusBadRequest)
		return
	}

	conAddress := chi.URLParam(r, "contractaddress")
	if conAddress == "" {
		app.ErrorJSON(w, errors.New("no contract address in url"))
	}

	smartContract, err := app.DB.GetSmartContract(conAddress)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	out, err := json.Marshal(smartContract)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("Content-Type", "application/json")
	w.Write(out)
}

func (app *Application) UpdateSmartContract(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPatch {
		app.ErrorJSON(w, errors.ErrUnsupported, http.StatusBadRequest)
		return
	}
	id := chi.URLParam(r, "contractaddress")
	var input models.SmartContract
	err := app.ReadJSON(w, r, &input)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	currentInDb, err := app.DB.GetSmartContract(id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	if input.Description != "" {
		currentInDb.Description = input.Description
	}
	if input.Abi != nil {
		currentInDb.Abi = input.Abi
	}
	if len(input.StateVariables) > 0 {
		currentInDb.StateVariables = input.StateVariables
	}
	if input.ProjectName != "" {
		currentInDb.ProjectName = input.ProjectName
	}

	// send call to update in db
	err = app.DB.UpdateSmartContractToAccountDb(id, *currentInDb)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	var successPayload = struct {
		Message       string `json:"message"`
		SmartContract models.SmartContract
	}{
		Message:       "Successfully updated contract",
		SmartContract: *currentInDb,
	}
	out, err := json.Marshal(successPayload)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("Content-Type", "application/json")
	w.Write(out)
}

// test @todo
func (app *Application) DeleteSmartContract(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		app.ErrorJSON(w, errors.ErrUnsupported, http.StatusBadRequest)
		return
	}
	connAdd := chi.URLParam(r, "contractaddress")
	userAddress := chi.URLParam(r, "address")
	err := app.DB.DeleteSmartContract(connAdd, userAddress)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, errors.ErrUnsupported, http.StatusBadRequest)
		return
	}

	var successPayload = struct {
		Message       string `json:"message"`
		SmartContract models.SmartContract
	}{
		Message: "Successfully deleted contract",
	}
	out, err := json.Marshal(successPayload)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("Content-Type", "application/json")
	w.Write(out)
}

func (app *Application) GetAllSmartContractAddressesByWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.ErrorJSON(w, errors.ErrUnsupported, http.StatusBadRequest)
		return
	}

	userId := chi.URLParam(r, "address")
	contractAddresses, err := app.DB.GetAllSmartContractInWalletAccounts(userId)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}
	out, err := json.Marshal(contractAddresses)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("Content-Type", "application/json")
	w.Write(out)

}

func (app *Application) GetSmartContractFullByWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.ErrorJSON(w, errors.ErrUnsupported, http.StatusBadRequest)
		return
	}

	userId := chi.URLParam(r, "address")
	contractAddresses, err := app.DB.GetAllFullScInWallet(userId)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}
	out, err := json.Marshal(contractAddresses)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("Content-Type", "application/json")
	w.Write(out)

}

func (app *Application) GetRemainCreditsByAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.ErrorJSON(w, errors.ErrUnsupported, http.StatusBadRequest)
		return
	}

	userId := chi.URLParam(r, "address")
	bal, err := app.Web3.GetRemainingCredits(userId)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	var payload = struct {
		Address string   `json:"address"`
		Balance *big.Int `json:"balance"`
	}{
		Address: userId,
		Balance: bal,
	}
	out, err := json.Marshal(payload)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("Content-Type", "application/json")
	w.Write(out)
}

func (app *Application) APIKeyWithLogin(w http.ResponseWriter, r *http.Request) {
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

	if input.Email == "" {
		app.ErrorJSON(w, errors.New("empty email field"))
		return
	}
	if input.Password == "" {
		app.ErrorJSON(w, errors.New("empty password field"))
		return
	}

	existing, err := app.DB.AdminGetWalletAccountByEmail(input.Email)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	passHash := app.HashPassword(input.Password)

	if existing.Password != passHash {
		app.ErrorJSON(w, errors.New("invalid login credentials"))
		return
	}

	out, err := json.Marshal(existing)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	// cookie := http.Cookie{
	// 	Name:    "apikey",
	// 	Value:   existing.ApiKey,
	// 	MaxAge:  86400000000000,
	// 	Expires: time.Now().Add(86400000000000),
	// }

	cookie := http.Cookie{
		Name:    "apikey",
		Value:   existing.ApiKey, // update value to have a signed token in order to validate on frontend
		MaxAge:  86400000000000,
		Expires: time.Now().Add(time.Duration(24 * time.Hour)),
		// SameSite: http.SameSiteStrictMode,
		Domain:   app.Domain,
		HttpOnly: true,
	}
	emaiCookie := http.Cookie{
		Name:    "email",
		Value:   existing.Email, // update value to have a signed token in order to validate on frontend
		MaxAge:  86400000000000,
		Expires: time.Now().Add(time.Duration(24 * time.Hour)),
		// SameSite: http.SameSiteStrictMode,
		Domain:   app.Domain,
		HttpOnly: true,
	}
	addressCookie := http.Cookie{
		Name:    "useraddress",
		Value:   existing.WalletAddress, // update value to have a signed token in order to validate on frontend
		MaxAge:  86400000000000,
		Expires: time.Now().Add(time.Duration(24 * time.Hour)),
		// SameSite: http.SameSiteStrictMode,
		Domain:   app.Domain,
		HttpOnly: true,
	}

	http.SetCookie(w, &addressCookie)
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &emaiCookie)


	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application")
	w.Write(out)

}

func (app *Application) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}
	id := chi.URLParam(r, "address")

	err := app.DB.ActivateAccount(id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	// VERIFY ACCOUNT IS ACTIVE BY RECALLING DB THEN RENDERING HTML WEBPAGE
	wallet, err := app.DB.AdminGetWalletAccount(id)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	if !wallet.Active {
		app.ErrorJSON(w, errors.New("error activating account"))
		return
	}

	var payload = struct {
		Message       string `json:"message"`
		WalletAddress string `json:"wallet_address"`
	}{
		Message:       "Successfully Activated Account",
		WalletAddress: wallet.WalletAddress,
	}
	out, err := json.Marshal(payload)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("Content-Type", "application/json")
	// render thank you for activating template

	w.Write(out)
}

// LOGIN HANDLER SETTING COOKIE IN BROWSER
func (app *Application) LoginWithEmail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

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

	validated := app.ValidateLogin(input)
	if !validated {
		app.ErrorJSON(w, errors.New("missing fields"))
		return
	}
	passHash := app.HashPassword(input.Password)

	existing, err := app.DB.AdminGetWalletAccountByEmail(input.Email)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	if existing.Password != passHash {
		app.ErrorJSON(w, errors.New("invalid login credentials"))
		return
	}

	out, err := json.Marshal(existing)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}


	cookie := http.Cookie{
		Name:    "apikey",
		Value:   existing.ApiKey, // update value to have a signed token in order to validate on frontend
		MaxAge:  86400000000000,
		Expires: time.Now().Add(time.Duration(24 * time.Hour)),
		// SameSite: http.SameSiteStrictMode,
		Domain:   app.Domain,
		HttpOnly: true,
	}
	emaiCookie := http.Cookie{
		Name:    "email",
		Value:   existing.Email, // update value to have a signed token in order to validate on frontend
		MaxAge:  86400000000000,
		Expires: time.Now().Add(time.Duration(24 * time.Hour)),
		// SameSite: http.SameSiteStrictMode,
		Domain:   app.Domain,
		HttpOnly: true,
	}
	addressCookie := http.Cookie{
		Name:    "useraddress",
		Value:   existing.WalletAddress, // update value to have a signed token in order to validate on frontend
		MaxAge:  86400000000000,
		Expires: time.Now().Add(time.Duration(24 * time.Hour)),
		// SameSite: http.SameSiteStrictMode,
		Domain:   app.Domain,
		HttpOnly: true,
	}

	http.SetCookie(w, &addressCookie)
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &emaiCookie)

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application")
	w.Write(out)

}

func (app *Application) ValidateApiKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}
	r.ParseForm()

	var input models.WalletAccount
	err := app.ReadJSON(w, r, &input)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}
	
	
	isValidApi, err := app.VerifyApiKeyHeader(input.ApiKey, input.Email)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}
	if !isValidApi {
		app.ErrorJSON(w, errors.New("apikey mismatch"))
		return
	}

	var payload = struct {
		Message string `json:"message"`
		Matches bool   `json:"matches"`
	}{
		Message: "api key valid",
		Matches: true,
	}
	out, err := json.Marshal(payload)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("Content-Type", "application/json")
	w.Write(out)
}


func(app *Application) Logout(w http.ResponseWriter, r *http.Request){
	cookie := http.Cookie{
		Name:    "apikey",
		Value:   "", // update value to have a signed token in order to validate on frontend
		MaxAge:  -1,
		Expires: time.Now(),
		// SameSite: http.SameSiteStrictMode,
		Domain:   app.Domain,
		HttpOnly: true,
	}
	emaiCookie := http.Cookie{
		Name:    "email",
		Value:   "", // update value to have a signed token in order to validate on frontend
		MaxAge:  -1,
		Expires: time.Now(),
		// SameSite: http.SameSiteStrictMode,
		Domain:   app.Domain,
		HttpOnly: true,
	}

	var payload = struct {
		Message string `json:"message"`
		LoggedOut bool `json:"logged_out"`

	}{
		Message: "successfully logged out",
		LoggedOut: true,
	}
	out, err := json.Marshal(payload)
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorJSON(w, err)
		return
	}

	http.SetCookie(w, &cookie)
	http.SetCookie(w, &emaiCookie)

	w.Header().Add("Content-Type", "application/json")
	w.Write(out)
}
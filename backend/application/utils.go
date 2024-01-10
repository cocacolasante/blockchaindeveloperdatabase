package application

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/models"
	"github.com/go-chi/chi/v5"
)

type JSONResponse struct {
	IsError bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (app *Application) writeJSON(w http.ResponseWriter, statusCode int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func (app *Application) ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()
	
	err := dec.Decode(data)
	if err != nil {
		app.ErrorLog.Println("Unexpected data in request body:", err)
		return err
	}

	var extraData struct{}
	err = dec.Decode(&extraData)
	if err != io.EOF {
		app.ErrorLog.Println("Unexpected data in request body:", err)
		return errors.New("body must only contain a single JSON value")
	}

	return nil

}

func (app *Application) ReadResponse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	body, err := io.ReadAll(r.Body)
	if err != nil {
		app.ErrorLog.Println("Error reading request body:", err)
		return err
	}

	app.InfoLog.Printf("Request Body: %s\n", body)
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()

	err = dec.Decode(data)
	if err != nil {
		app.ErrorLog.Println("Error decoding JSON:", err)
		return err
	}

	return nil
}

func (app *Application) ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.IsError = true
	payload.Message = err.Error()

	return app.writeJSON(w, statusCode, payload)
}

func (app *Application) HashPassword(password string) string {
	passHash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(passHash[:])
}

func (app *Application) ValidateSignUp(input *models.WalletAccount) bool {
	if input.Email == "" || input.Password == "" {
		return false
	}

	return true
}

func (app *Application) VerifyHeaders(w http.ResponseWriter, r *http.Request) (bool, error) {
	isVerified := false
	w.Header().Add("Vary", "Authorization")

	// get auth header
	authHeader := r.Header.Get("Authorization")

	// sanity check
	if authHeader == "" {
		return isVerified, errors.New("no auth header")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return isVerified, errors.New("invalid auth header")
	}

	// check to see if we have the word Bearer
	if headerParts[0] != "Bearer" {
		return isVerified, errors.New("invalid auth header")
	}

	token := headerParts[1]

	id := chi.URLParam(r, "address")
	
	matches, err := app.CheckIfApiMatchesDatabase(id, token)
	if err != nil {
		return isVerified, err
	}
	if matches{
		isVerified = true
	}

	return isVerified, nil
}


func(app *Application) CheckIfApiMatchesDatabase(walletAddress, key string) (bool, error) {
	wallet, err := app.DB.AdminGetWalletAccount(walletAddress) 
	
	if err != nil {
		return false, err
	}
	if key != wallet.ApiKey {
		return false, errors.New("invalid auth credentials")
	}
	return true, nil
}



func(app *Application) VerifyURL(r *http.Request) (bool, error) {
	walletAddress := chi.URLParam(r, "address")
	reqKey := r.URL.Query().Get("key")
	
	wallet, err := app.DB.AdminGetWalletAccount(walletAddress) 
		
	if err != nil {
		return false, err
	}
	

	if wallet.ApiKey != reqKey {
		return false, errors.New("invalid email auth")
	}
	
	
	return true, nil

}


func(app *Application) VerifyActive(r *http.Request) (bool, error){
	walletAddress := chi.URLParam(r, "address")
	wallet, err := app.DB.AdminGetWalletAccount(walletAddress) 
	if err != nil {
		return false, err
	}

	if !wallet.Active{
		return false, errors.New("inactive account")
	}

	return true, nil
}
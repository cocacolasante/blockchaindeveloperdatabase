package application

import (
	"errors"
	"net/http"
)

func (app *Application) EnableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://*")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")

			return

		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func (app *Application) authRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isVerified, err := app.VerifyHeaders(w, r)
		if err != nil {
			app.ErrorJSON(w, err)
			return
		}

		if !isVerified {
			app.ErrorJSON(w, errors.New("unauthenticated"), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *Application) EmailAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		verified, err := app.VerifyURL(r)
		if err != nil {
			app.ErrorJSON(w, err)
			return
		}

		if !verified {
			app.ErrorJSON(w, errors.New("unauthenticated email"), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func(app *Application) ActiveAccountMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		valid, err := app.VerifyActive(r)
		if err != nil {
			app.ErrorJSON(w, err)
			return
		}

		if !valid {
			app.ErrorJSON(w, errors.New("inactive account"), http.StatusBadRequest)
			return
		}
		

		next.ServeHTTP(w, r)
	})
}


func(app *Application) AdminAuthentication(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		

		next.ServeHTTP(w, r)
	})
}
package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
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
	app.InfoLog.Println("hit in read json")
	err := dec.Decode(data)
	if err != nil {
		app.ErrorLog.Println("Unexpected data in request body:", err)
		return err
	}
	
	var extraData struct{}
	err = dec.Decode(&extraData)
	if err != io.EOF {
		app.ErrorLog.Println("Unexpected data in request body:", err)
		return errors.New("Body must only contain a single JSON value")
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

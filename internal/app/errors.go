package app

import (
	"fmt"
	"net/http"
)

// LogError logs an error message.
func (app *App) LogError(r *http.Request, err error) {
	app.Logger.Errorf("%s: %s", r.URL.Path, err)
}

// ErrorResponse sends a JSON response with a given status code and message.
func (app *App) ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	errorResponse := Envelope{
		"error": message,
	}

	err := app.WriteJSON(w, status, errorResponse, nil)
	if err != nil {
		app.LogError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ServerErrorResponse sends a JSON response with a StatusInternalServerError status code and message.
func (app *App) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.LogError(r, err)

	message := "the server encountered a problem and could not process your request."
	app.ErrorResponse(w, r, http.StatusInternalServerError, message)
}

// NotFoundResponse sends a JSON response with a StatusNotFound status code and message.
func (app *App) NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found."
	app.ErrorResponse(w, r, http.StatusNotFound, message)
}

// MethodNotAllowedResponse sends a JSON response with a StatusMethodNotAllowed status code and message.
func (app *App) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource.", r.Method)
	app.ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// FailedValidationResponse sends a JSON response with a StatusUnprocessableEntity status code and messages.
func (app *App) FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.ErrorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

package app

import (
	"net/http"
)

func (app *App) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	response := Envelope{
		"status": "ok",
	}

	err := app.WriteJSON(w, http.StatusOK, response, nil)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
	}
}

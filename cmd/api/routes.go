package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/accounts", app.showAListOfAccountsHandler)

	router.HandlerFunc(http.MethodGet, "/v1/accounts/:id", app.showAnAccountHandler)

	router.HandlerFunc(http.MethodPatch, "/v1/accounts/:id", app.updateAnAccountHandler)

	router.HandlerFunc(http.MethodPost, "/v1/accounts", app.createAnAccountHandler)

	return router
}

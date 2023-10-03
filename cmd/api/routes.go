package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/accounts", app.showAListOfAccountsHandler)

	router.HandlerFunc(http.MethodGet, "/v1/accounts/:id", app.showAnAccountHandler)

	router.HandlerFunc(http.MethodPatch, "/v1/accounts/:id", app.updateAnAccountHandler)

	router.HandlerFunc(http.MethodPost, "/v1/accounts", app.createAnAccountHandler)

	router.HandlerFunc(http.MethodDelete, "/v1/accounts/:id", app.deleteAnAccountHandler)

	return app.recoverPanic(app.rateLimit(router))
}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/saleh-ghazimoradi/bank-system.git/internal/data"
)

func (app *application) showAListOfAccountsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "show a lists of accounts")
}

func (app *application) showAnAccountHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	account := data.Account{
		ID:        id,
		CreatedAt: time.Now(),
		FirstName: "Saleh",
		LastName:  "Ghazimoradi",
		Number:    int64(rand.Intn(9999999999999999)),
		Balance:   2500,
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)

	if err != nil {
		app.logger.Error(err.Error())
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createAnAccountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create an account")
}

func (app *application) deleteAnAccountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete an account by ID")
}

func (app *application) updateAnAccountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update an account")
}

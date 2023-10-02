package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/saleh-ghazimoradi/bank-system.git/internal/data"
	"github.com/saleh-ghazimoradi/bank-system.git/internal/validator"
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
	var input struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	account := &data.Account{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	v := validator.New()

	if data.ValidateAccount(v, account); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) deleteAnAccountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete an account by ID")
}

func (app *application) updateAnAccountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update an account")
}

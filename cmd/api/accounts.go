package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"

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

	account, err := app.models.Account.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
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
		Balance   int64  `json:"balance"`
		Number    int64  `json:"number"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	account := &data.Account{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Balance:   input.Balance,
		Number:    int64(rand.Intn(999999999999)),
	}

	v := validator.New()

	if data.ValidateAccount(v, account); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Account.Insert(account)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/accounts/%d", account.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"account": account}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteAnAccountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete an account by ID")
}

func (app *application) updateAnAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	account, err := app.models.Account.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Balance   int64  `json:"balance"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	account.FirstName = input.FirstName
	account.LastName = input.LastName
	account.Balance = input.Balance

	v := validator.New()

	if data.ValidateAccount(v, account); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Account.Update(account)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

func (app *application) showAListOfAccountsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "show a lists of accounts")
}

func (app *application) showAnAccountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "show an account by ID")
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

func (app *application) transferMoneyToAnAccountHandler(w http.ResponseWriter, r *http.ResponseWriter) {
	fmt.Fprintln(w, "transfer money to an account")
}

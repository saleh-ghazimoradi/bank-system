package data

import (
	"database/sql"
	"time"

	"github.com/saleh-ghazimoradi/bank-system.git/internal/validator"
)

type Account struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	Version   int32     `json:"version"`
}

func ValidateAccount(v *validator.Validator, account *Account) {
	v.Check(account.FirstName != "", "firstName", "must be provided")
	v.Check(len(account.FirstName) >= 3 && len(account.FirstName) <= 15, "firstName", "must neither be less than 3 nor greater than 15 bytes long")

	v.Check(account.LastName != "", "last name", "must be provided")
	v.Check(len(account.LastName) >= 3 && len(account.LastName) <= 15, "lastName", "must neither be less than 3 nor greater than 15 bytes long")

}

type AccountModel struct {
	DB *sql.DB
}

func (a AccountModel) Insert(account *Account) error {
	return nil
}

func (a AccountModel) Get(id int64) (*Account, error) {
	return nil, nil
}

func (a AccountModel) Update(account *Account) error {
	return nil
}

func (a AccountModel) Delete(id int64) error {
	return nil
}

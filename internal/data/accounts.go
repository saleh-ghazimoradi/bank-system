package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/saleh-ghazimoradi/bank-system.git/internal/validator"
)

type Account struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Balance   int64     `json:"balance"`
	Number    int64     `json:"number"`
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
	query := `
        INSERT INTO bank (first_name, last_name, balance, number) 
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, version`

	args := []any{account.FirstName, account.LastName, account.Balance, account.Number}
	return a.DB.QueryRow(query, args...).Scan(&account.ID, &account.CreatedAt, &account.Version)
}

func (a AccountModel) Get(id int64) (*Account, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT id, created_at, first_name, last_name, balance, number, version
	FROM bank
	WHERE id = $1`

	var account Account

	err := a.DB.QueryRow(query, id).Scan(
		&account.ID,
		&account.CreatedAt,
		&account.FirstName,
		&account.LastName,
		&account.Balance,
		&account.Number,
		&account.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &account, nil

}

func (a AccountModel) Update(account *Account) error {
	query := `
        UPDATE bank
        SET first_name = $1, last_name = $2, balance = $3, version = version + 1
        WHERE id = $4
        RETURNING version`

	args := []any{
		account.FirstName,
		account.LastName,
		account.Balance,
		account.ID,
	}

	return a.DB.QueryRow(query, args...).Scan(&account.Version)
}

func (a AccountModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `DELETE FROM bank
        WHERE id = $1`

	result, err := a.DB.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

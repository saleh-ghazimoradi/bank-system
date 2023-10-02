package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Account AccountModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Account: AccountModel{DB: db},
	}
}

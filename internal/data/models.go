package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Account AccountModel
	Users   UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Account: AccountModel{DB: db},
		Users:   UserModel{DB: db},
	}
}

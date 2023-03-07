package infra

import (
	"database/sql"
)

type CurrencyRepository struct {
	db *sql.DB
}

func NewCurrencyRepository(db *sql.DB) CurrencyRepository {
	return CurrencyRepository{
		db: db,
	}
}

package infra

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ArtuoS/hurb/internal/domain/entity"
)

type CurrencyRepository struct {
	db *sql.DB
}

func NewCurrencyRepository(db *sql.DB) CurrencyRepository {
	return CurrencyRepository{
		db: db,
	}
}

func (c *CurrencyRepository) Get(expression string) ([]entity.Currency, error) {
	var currencies []entity.Currency

	query := fmt.Sprintf("SELECT * FROM currencies WHERE %s", expression)
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, errors.New("failed to query currencies from database")
	}

	for rows.Next() {
		var currency entity.Currency
		err := rows.Scan(
			&currency.Id,
			&currency.Code,
			&currency.Symbol,
			&currency.DollarValue,
		)
		if err != nil {
			return nil, errors.New("failed to scan row into currency entity")
		}
		currencies = append(currencies, currency)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New("failed to iterate over currency result rows")
	}

	return currencies, nil
}

package entity

import "github.com/google/uuid"

type Currency struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	DollarValue float64   `json:"dollar_value"`
	Symbol      string    `json:"symbol"`
}

func NewCurrency(name, symbol string, dollarValue float64) *Currency {
	return &Currency{
		Id:          uuid.New(),
		Name:        name,
		DollarValue: dollarValue,
		Symbol:      symbol,
	}
}

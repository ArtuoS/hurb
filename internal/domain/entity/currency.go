package entity

import "github.com/google/uuid"

type Currency struct {
	Id          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	DollarValue float32   `json:"dollar_value"`
	Symbol      string    `json:"symbol"`
}

func NewCurrency(code, description, symbol string, dollarValue float32) *Currency {
	return &Currency{
		Id:          uuid.New(),
		Code:        code,
		Description: description,
		DollarValue: dollarValue,
		Symbol:      symbol,
	}
}

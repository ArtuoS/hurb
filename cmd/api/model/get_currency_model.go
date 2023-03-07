package model

import "github.com/google/uuid"

type GetCurrencyModel struct {
	From   uuid.UUID `json:"from"`
	To     uuid.UUID `json:"to"`
	Amount float64   `json:"amount"`
}

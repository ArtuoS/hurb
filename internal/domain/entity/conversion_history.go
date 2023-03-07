package entity

import (
	"time"

	"github.com/google/uuid"
)

type ConversionHistory struct {
	From uuid.UUID `json:"from"`
	To   uuid.UUID `json:"to"`
	Date time.Time `json:"date"`
}

func NewConversionHistory(from, to uuid.UUID) *ConversionHistory {
	return &ConversionHistory{
		From: from,
		To:   to,
		Date: time.Now(),
	}
}

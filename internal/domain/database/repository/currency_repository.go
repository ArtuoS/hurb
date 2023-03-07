package repository

import (
	"fmt"

	"github.com/ArtuoS/hurb/internal/domain/entity"
	"github.com/google/uuid"
)

type CurrencyRepositoryInterface interface {
	Create(entity.Currency) error
	GetAll() ([]entity.Currency, error)
	GetById(uuid.UUID) (entity.Currency, error)
	GetByName(string) (entity.Currency, error)
}

type CurrencyRepository struct {
}

func (c *CurrencyRepository) Create(entity.Currency) error {
	fmt.Println("creating currency")
	return nil
}

func (c *CurrencyRepository) GetByName(string) (entity.Currency, error) {
	fmt.Println("creating currency")
	return nil
}

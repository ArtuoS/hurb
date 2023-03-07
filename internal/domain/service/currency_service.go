package service

import "github.com/ArtuoS/hurb/internal/domain/database/repository"

type CurrencyService struct {
	repository repository.CurrencyRepository
}

func NewCurrencyService(currencyRepository repository.CurrencyRepository) CurrencyService {
	return CurrencyService{
		repository: currencyRepository,
	}
}

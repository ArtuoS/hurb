package service

import (
	"fmt"

	"github.com/ArtuoS/hurb/internal/domain/database/infra"
	"github.com/ArtuoS/hurb/pb"
)

type CurrencyService struct {
	repository infra.CurrencyRepository
}

func NewCurrencyService(currencyRepository infra.CurrencyRepository) CurrencyService {
	return CurrencyService{
		repository: currencyRepository,
	}
}

func (c *CurrencyService) ConvertCurrency(model *pb.ConvertCurrencyRequest) (pb.ConvertCurrencyResponse, error) {
	fmt.Println("Received in switch: ", model.GetFrom())
	return pb.ConvertCurrencyResponse{}, nil
}

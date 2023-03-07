package service

import (
	"errors"
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
	expression := fmt.Sprintf("code in (%s, %s)", model.GetFrom(), model.GetTo())
	result, err := c.repository.Get(expression)
	if err != nil {
		return pb.ConvertCurrencyResponse{}, err
	}

	if len(result) != 2 {
		return pb.ConvertCurrencyResponse{}, errors.New("currencies not found")
	}

	return pb.ConvertCurrencyResponse{
		Amount: func() float32 {
			return (result[1].DollarValue * model.GetAmount()) / result[0].DollarValue
		}(),
	}, nil
}

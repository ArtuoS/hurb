package server

import (
	"context"

	"github.com/ArtuoS/hurb/internal/domain/database/infra"
	"github.com/ArtuoS/hurb/internal/domain/service"
	"github.com/ArtuoS/hurb/pb"
)

type Server struct {
	pb.UnimplementedCurrencyServer
}

func (s *Server) ConvertCurrency(ctx context.Context, in *pb.ConvertCurrencyRequest) (*pb.ConvertCurrencyResponse, error) {
	service := service.NewCurrencyService(
		infra.NewCurrencyRepository(nil),
	)

	convertedValue, err := service.ConvertCurrency(in)
	if err != nil {
		return nil, err
	}

	return &convertedValue, nil
}

package server

import (
	"context"
	"github/nergilz/taskGetRate/internal/domain"

	ratev1 "github.com/nergilz/grpcTaskGetRate/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// интерфейс get rates usdt из сервисного слоя
type IGetRates interface {
	GetRates(ctx context.Context, market string) (domain.RateResponse, error)
}

// реализует функционал api
type serverAPI struct {
	ratev1.UnimplementedRatesServer
	api IGetRates
}

func Register(grpcs *grpc.Server, getrate IGetRates) {
	ratev1.RegisterRatesServer(grpcs, &serverAPI{api: getrate})
}

func (s *serverAPI) GetRates(ctx context.Context, in *ratev1.RateRequest) (*ratev1.RateResponse, error) {
	if in.Market == "" {
		return nil, status.Error(codes.InvalidArgument, "market is required")
	}

	if in.Market != "usdt" {
		return nil, status.Errorf(codes.InvalidArgument, "can not request with market %s", in.Market)
	}

	resp, err := s.api.GetRates(ctx, in.GetMarket())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "can not get rate")
	}

	return &ratev1.RateResponse{
		Market:    resp.Market,
		Timestamp: resp.TimeStamp,
		Ask: &ratev1.Data{
			Price:   resp.Ask.Price,
			Volume:  resp.Ask.Volume,
			Amount:  resp.Ask.Amount,
			Factor:  resp.Ask.Factor,
			TypeAsk: resp.Ask.TypeAsk,
		},
		Bid: &ratev1.Data{
			Price:   resp.Bid.Price,
			Volume:  resp.Bid.Volume,
			Amount:  resp.Bid.Amount,
			Factor:  resp.Bid.Factor,
			TypeAsk: resp.Bid.TypeAsk,
		},
	}, nil
}

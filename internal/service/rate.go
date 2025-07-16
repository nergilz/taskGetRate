package service

import (
	"context"
	"github/nergilz/taskGetRate/internal/domain"
	"log/slog"
)

type IMarketStorage interface {
	SetMarket(ctx context.Context, markerRate domain.RateResponse) error
}

type IMarketTransport interface {
	GetDataFromGrinexApi(ctx context.Context, market string) (domain.RateResponse, error)
}

type Rate struct {
	logger    *slog.Logger
	store     IMarketStorage
	transport IMarketTransport
}

func New(log *slog.Logger, store IMarketStorage, transport IMarketTransport) *Rate {
	return &Rate{
		logger:    log,
		store:     store,
		transport: transport,
	}
}

func (r *Rate) GetRates(ctx context.Context, market string) (domain.RateResponse, error) {
	resp, err := r.transport.GetDataFromGrinexApi(ctx, market)
	if err != nil {
		r.logger.Error("failed get rates", domain.Err(err))

		return domain.RateResponse{}, err
	}

	err = r.store.SetMarket(ctx, resp)
	if err != nil {
		r.logger.Error("failed save market rates", domain.Err(err))

		return domain.RateResponse{}, err
	}

	return resp, nil
}

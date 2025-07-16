package storage

import (
	"context"
	"github/nergilz/taskGetRate/internal/domain"
)

type AppStorage struct {
}

func New() AppStorage {
	return AppStorage{}
}

func (s AppStorage) SetMarket(ctx context.Context, markerRate domain.RateResponse) error {

	return nil
}

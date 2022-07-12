package currency

import (
	"context"

	"github.com/IlmarLopez/currency/internal/entity"
	"go.uber.org/zap"
)

// Service is the interface that provides the currency service.
type Service interface {
	Count(ctx context.Context) (int, error)
	Query(ctx context.Context, offset, limit int) ([]Currency, error)
}

// service is the implementation of the Service interface.
type service struct {
	repo   Repository
	logger *zap.SugaredLogger
}

// NewService returns a new currency service.
func NewService(repo Repository, logger *zap.SugaredLogger) Service {
	return service{repo, logger}
}

// Currency represents the data about a currency record.
type Currency struct {
	entity.Currency
}

// Count returns the number of currencies.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the list of currencies with the given offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Currency, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Currency{}
	for _, item := range items {
		result = append(result, Currency{item})
	}
	return result, nil
}

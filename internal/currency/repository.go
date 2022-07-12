package currency

import (
	"context"

	"github.com/IlmarLopez/currency/internal/entity"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

// Repository is the interface that provides the currency repository.
type Repository interface {
	// Count returns the number of currencies.
	Count(ctx context.Context, currency string, queryParameters map[string]string) (int, error)
	// Query returns the list of currencies with the given offset and limit.
	Query(ctx context.Context, currency string, queryParameters map[string]string, offset, limit int) ([]entity.Currency, error)
}

// repository is the implementation of the Repository interface.
type repository struct {
	db     *pgx.Conn
	logger *zap.SugaredLogger
}

// NewRepository returns a new currency repository.
func NewRepository(db *pgx.Conn, logger *zap.SugaredLogger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

// Count returns the number of currencies.
func (r repository) Count(ctx context.Context, currency string, queryParameters map[string]string) (int, error) {
	var count int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM currencies").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Query returns the list of currencies with the given offset and limit.
func (r repository) Query(ctx context.Context, currency string, queryParameters map[string]string, offset, limit int) ([]entity.Currency, error) {
	var currencies []entity.Currency
	rows, err := r.db.Query(ctx, "SELECT * FROM currencies LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currency entity.Currency
		err := rows.Scan(&currency.ID)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

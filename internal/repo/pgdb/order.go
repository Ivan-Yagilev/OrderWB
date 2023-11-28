package pgdb

import (
	"context"
	"order/internal/entity"
	"order/pkg/postgres"
)

type OrderRepo struct {
	*postgres.Postgres
}

func NewOrderRepo(pg *postgres.Postgres) *OrderRepo {
	return &OrderRepo{pg}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, input entity.Order) error {
	// Переделать базу (orders|items)

	// Кӕширование

	// sqlx or pgx?

	return nil
}

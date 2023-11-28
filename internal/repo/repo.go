package repo

import (
	"context"
	"order/internal/entity"
	"order/internal/repo/pgdb"
	"order/pkg/postgres"
)

type Order interface {
	CreateOrder(ctx context.Context, input entity.Order) error
}

type Repositories struct {
	Order
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Order: pgdb.NewOrderRepo(pg),
	}
}

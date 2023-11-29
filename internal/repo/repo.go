package repo

import (
	"context"
	"order/internal/entity"
	"order/internal/repo/pgdb"
	"order/pkg/postgres"
)

type OrderPostgres interface {
	CreateOrder(ctx context.Context, input entity.Order) error
}

type OrderCache interface {
}

type Repositories struct {
	OrderPostgres
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		OrderPostgres: pgdb.NewOrderRepo(pg),
	}
}

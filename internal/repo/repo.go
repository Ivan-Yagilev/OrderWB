package repo

import (
	"order/internal/repo/pgdb"
	"order/pkg/postgres"
)

type Something interface {
}

type Repositories struct {
	Something
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Something: pgdb.NewSomethingRepo(pg),
	}
}

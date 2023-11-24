package pgdb

import "order/pkg/postgres"

type SomethingRepo struct {
	*postgres.Postgres
}

func NewSomethingRepo(pg *postgres.Postgres) *SomethingRepo {
	return &SomethingRepo{pg}
}

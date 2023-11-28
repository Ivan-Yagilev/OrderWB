package service

import (
	"context"
	"order/internal/entity"
	"order/internal/repo"
	"order/pkg/hasher"
	"time"
)

type Order interface {
	CreateOrder(ctx context.Context, input entity.Order) error
}

type Services struct {
	Order Order
}

type ServicesDependencies struct {
	Repos  *repo.Repositories
	Hasher hasher.PasswordHasher

	SignKey  string
	TokenTTL time.Duration
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Order: NewOrderService(deps.Repos.Order),
	}
}

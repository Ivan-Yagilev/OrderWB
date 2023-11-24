package service

import (
	"order/internal/repo"
	"order/pkg/hasher"
	"time"
)

type Services struct {
}

type ServicesDependencies struct {
	Repos  *repo.Repositories
	Hasher hasher.PasswordHasher

	SignKey  string
	TokenTTL time.Duration
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{}
}

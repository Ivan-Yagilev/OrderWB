package service

import (
	"order/internal/repo"
	"order/pkg/hasher"
	"time"
)

type Some interface {
}

type Services struct {
	Some Some
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

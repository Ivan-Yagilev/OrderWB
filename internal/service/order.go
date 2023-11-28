package service

import (
	"context"
	"order/internal/entity"
	"order/internal/repo"
	"order/internal/repo/repoerrs"
)

type OrderService struct {
	orderRepo repo.Order
}

func NewOrderService(orderRepo repo.Order) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, input entity.Order) error {
	err := s.orderRepo.CreateOrder(ctx, input)
	if err != nil {
		if err == repoerrs.ErrAlreadyExists {
			return ErrOrderAlreadyExists
		}
		return ErrCannotCreateOrder
	}

	return nil
}

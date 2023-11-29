package pgdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"order/internal/entity"
	"order/internal/repo/repoerrs"
	"order/pkg/postgres"
)

type OrderRepo struct {
	*postgres.Postgres
}

func NewOrderRepo(pg *postgres.Postgres) *OrderRepo {
	return &OrderRepo{pg}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, input entity.Order) error {
	items := input.Items

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("OrderRepo.CreateOrder - r.Pool.Begin: %v", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	sql, args, _ := r.Builder.
		Insert("orders").
		Columns(
			"order_uid", "track_number", "entry",
			"name", "phone", "zip", "city", "address", "region", "email",
			"transaction", "request_id", "currency", "provider", "amount", "payment_dt", "bank", "delivery_cost", "goods_total", "custom_fee",
			"locale", "internal_signature", "customer_id", "delivery_service", "shardkey", "sm_id", "date_created", "oof_shard",
		).
		Values(
			input.OrderUid, input.TrackNumber, input.Entry,
			//
			input.Delivery.Name, input.Delivery.Phone, input.Delivery.Zip, input.Delivery.City,
			input.Delivery.Address, input.Delivery.Region, input.Delivery.Email,
			//
			input.Payment.Transaction, input.Payment.RequestId, input.Payment.Currency, input.Payment.Provider,
			input.Payment.Amount, input.Payment.PaymentDt, input.Payment.Bank, input.Payment.DeliveryCost,
			input.Payment.GoodsTotal, input.Payment.CustomFee,
			//
			input.Locale, input.InternalSignature, input.CustomerId, input.DeliveryService, input.Shardkey, input.SmId,
			input.DateCreated, input.OofShard,
		).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return repoerrs.ErrAlreadyExists
			}
		}
		return fmt.Errorf("OrderRepo.CreateOrder - tx.Exec: %v", err)
	}

	for _, item := range items {
		sql, args, _ := r.Builder.
			Insert("items").
			Columns(
				"order_id", "chrt_id", "track_number", "price", "rid",
				"name", "sale", "size", "total_price", "nm_id", "brand", "status",
			).
			Values(
				input.OrderUid,
				item.ChrtId, item.TrackNumber, item.Price, item.Rid,
				item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status,
			).
			ToSql()

		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			return fmt.Errorf("OrderRepo.CreateItem - tx.Exec: %v", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("OrderRepo.CreateOrder - tx.Commit: %v", err)
	}

	return nil

	// Переделать базу (orders|items)

	// Кӕширование

	// sqlx or pgx?
}

package services

import (
	"context"
	"errors"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"go.uber.org/zap"
)

type OrdersRepo interface {
	CreateOrder(ctx context.Context, order *models.Order) (string, error)
	GetOrderById(ctx context.Context, id string) (*models.Order, error)
	GetUserOrders(ctx context.Context, userId string) ([]*models.Order, error)
	GetOrders(ctx context.Context, query types.OrderFilters) ([]*models.Order, error)
	UpdateOrder(ctx context.Context, order *models.Order) error
	DeleteOrder(ctx context.Context, id string) error
}

type OrderService struct {
	log      *zap.Logger
	repo     OrdersRepo
	userRepo UserRepo
}

func NewOrderService(
	logger *zap.Logger,
	repo OrdersRepo,
	userRepo UserRepo,
) *OrderService {
	return &OrderService{
		log:      logger,
		repo:     repo,
		userRepo: userRepo,
	}
}

func (r *OrderService) CreateOrder(ctx context.Context, order *models.Order) (string, error) {
	l := r.log.With(
		zap.Any("operation", "OrderService.CreateOrder"),
		zap.Any("userId", order.UserID),
	)

	id, err := r.repo.CreateOrder(ctx, order)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyExist) {
			l.Info("order already exist")
			return "", ErrOrderAlreadyExist
		}
		l.Error(
			"failed to create order",
			zap.Any("error", err),
		)
		return "", err
	}
	return id, nil
}

func (r *OrderService) GetOrderById(ctx context.Context, id string) (*models.Order, error) {
	l := r.log.With(
		zap.Any("operation", "OrderService.GetOrderById"),
		zap.Any("orderId", id),
	)

	order, err := r.repo.GetOrderById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("order not found")
			return nil, ErrOrderNotFound
		}
		l.Error(
			"failed to get order",
			zap.Any("error", err),
		)
		return nil, err
	}
	return order, nil
}

func (r *OrderService) GetUserOrders(ctx context.Context, userID string) ([]*models.Order, error) {
	l := r.log.With(
		zap.Any("operation", "OrderService.GetUserOrders"),
		zap.Any("userID", userID),
	)

	orders, err := r.repo.GetUserOrders(ctx, userID)
	if err != nil {
		l.Error(
			"failed to get user orders",
			zap.Any("error", err),
		)
		return nil, err
	}
	return orders, nil
}

func (r *OrderService) GetOrders(ctx context.Context, query types.OrderFilters) ([]*models.Order, error) {
	l := r.log.With(
		zap.Any("operation", "OrderService.GetOrders"),
	)

	if query.Worker.Email != "" ||
		query.Worker.Name != "" ||
		query.Worker.Surname != "" ||
		query.Worker.Patronymic != "" {
		worker, err := r.userRepo.GetUsers(ctx, types.UserFilters{
			UserData: query.Worker,
			UserType: "WORKER",
		})
		if err != nil {
			l.Error(
				"failed to get workers",
				zap.Any("error", err),
			)
			return nil, err
		}
		for _, w := range worker {
			query.WorkersID = append(query.WorkersID, w.ID.Hex())
		}
	}

	l.Info("query", zap.Any("query", query))
	orders, err := r.repo.GetOrders(ctx, query)
	if err != nil {
		l.Error(
			"failed to get orders",
			zap.Any("error", err),
		)
		return nil, err
	}
	l.Info("orders", zap.Any("orders", orders))

	return orders, nil
}

func (r *OrderService) UpdateOrder(ctx context.Context, order *models.Order) error {
	l := r.log.With(
		zap.Any("operation", "OrderService.UpdateOrder"),
		zap.Any("orderID", order.ID),
	)

	if err := r.repo.UpdateOrder(ctx, order); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("order not found")
			return ErrOrderNotFound
		}
		l.Error(
			"failed to update order",
			zap.Any("error", err),
		)
		return err
	}
	return nil
}

func (r *OrderService) DeleteOrder(ctx context.Context, id string) error {
	l := r.log.With(
		zap.Any("operation", "OrderService.DeleteOrder"),
		zap.Any("orderID", id),
	)

	if err := r.repo.DeleteOrder(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("order not found")
			return ErrOrderNotFound
		}
		l.Error(
			"failed to delete order",
			zap.Any("error", err),
		)
		return err
	}
	return nil
}

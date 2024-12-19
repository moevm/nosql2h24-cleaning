package services

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"go.mongodb.org/mongo-driver/v2/bson"
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

type WorkerRepo interface {
	IncrementWorkersOrderCounts(ctx context.Context, workerID ...string) error
}

type OrderService struct {
	log        *zap.Logger
	repo       OrdersRepo
	userRepo   UserRepo
	workerRepo WorkerRepo
}

func NewOrderService(
	logger *zap.Logger,
	repo OrdersRepo,
	userRepo UserRepo,
	workerRepo WorkerRepo,
) *OrderService {
	return &OrderService{
		log:        logger,
		repo:       repo,
		userRepo:   userRepo,
		workerRepo: workerRepo,
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

	orders, err := r.repo.GetOrders(ctx, query)
	if err != nil {
		l.Error(
			"failed to get orders",
			zap.Any("error", err),
		)
		return nil, err
	}

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

func (r *OrderService) AssignWorkerToOrder(ctx context.Context, orderID, workerID string) error {
	l := r.log.With(
		zap.Any("operation", "OrderService.AssignWorkerToOrder"),
		zap.Any("orderID", orderID),
		zap.Any("workerID", workerID),
	)

	order, err := r.GetOrderById(ctx, orderID)
	if err != nil {
		l.Error(
			"failed to get order",
			zap.Any("error", err),
		)
		return err
	}

	if order.Status == models.OrderStatusDone {
		l.Info("order already done")
		return ErrOrderAlreadyDone
	}

	worker, err := r.userRepo.GetUserById(ctx, workerID)
	if err != nil {
		l.Error(
			"failed to get worker",
			zap.Any("error", err),
		)
		return err
	}

	if slices.Contains(order.Workers, worker.ID) {
		l.Info("worker already assigned")
		return ErrWorkerAlreadyAssigned
	}

	if order.Status == models.OrderStatusNew {
		order.StatusLogs = append(order.StatusLogs, models.StatusLog{
			ID:         bson.NewObjectID(),
			PrevStatus: order.Status,
			NewStatus:  models.OrderStatusAccepted,
			CreatedAt:  time.Now(),
		})
		order.Status = models.OrderStatusAccepted
	}

	order.Workers = append(order.Workers, worker.ID)

	if err := r.UpdateOrder(ctx, order); err != nil {
		l.Error(
			"failed to update order",
			zap.Any("error", err),
		)
		return err
	}

	return nil
}

func (r *OrderService) CompleteOrder(ctx context.Context, orderID string) error {
	l := r.log.With(
		zap.Any("operation", "OrderService.CompleteOrder"),
		zap.Any("orderID", orderID),
	)

	order, err := r.GetOrderById(ctx, orderID)
	if err != nil {
		l.Error(
			"failed to get order",
			zap.Any("error", err),
		)
		return err
	}

	if order.Status != models.OrderStatusAccepted {
		l.Info("order not done")
		return ErrOrderIncorrectStatus
	}

	order.StatusLogs = append(order.StatusLogs, models.StatusLog{
		ID:         bson.NewObjectID(),
		PrevStatus: order.Status,
		NewStatus:  models.OrderStatusDone,
		CreatedAt:  time.Now(),
	})
	order.Status = models.OrderStatusDone

	if err := r.UpdateOrder(ctx, order); err != nil {
		l.Error(
			"failed to update order",
			zap.Any("error", err),
		)
		return err
	}
	
	ids := make([]string, len(order.Workers))
	for i, id := range order.Workers {
		ids[i] = id.Hex()
	}

	if err := r.workerRepo.IncrementWorkersOrderCounts(ctx, ids...); err != nil {
		l.Error(
			"failed to increment workers order counts",
			zap.Any("error", err),
		)
		return err
	}

	return nil
}

package services

import (
	"context"
	"errors"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

type ServiceRepo interface {
	CreateService(ctx context.Context, service *models.Service) (string, error)
	GetService(ctx context.Context, id string) (*models.Service, error)
	GetServices(ctx context.Context) ([]*models.Service, error)
	UpdateService(ctx context.Context, service *models.Service) error
	DeleteService(ctx context.Context, id string) error
}

// Bad naming
type Service struct {
	log  *zap.Logger
	repo ServiceRepo
}

func NewService(
	logger *zap.Logger,
	repo ServiceRepo,
) *Service {
	return &Service{
		log:  logger,
		repo: repo,
	}
}

func (r *Service) CreateService(ctx context.Context, service *models.Service) (string, error) {
	l := r.log.With(
		zap.Any("operation", "Service.CreateService"),
	)

	// TODO: make new collection?
	for i := 0; i < len(service.Consumables); i++ {
		service.Consumables[i].ID = bson.NewObjectID()
	}

	id, err := r.repo.CreateService(ctx, service)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyExist) {
			l.Info("service already exist", zap.Any("error", err))
			return "", ErrServiceAlreadyExist
		}
		l.Error("creation failed", zap.Any("error", err))
		return "", err
	}

	return id, nil
}

func (r *Service) GetService(ctx context.Context, id string) (*models.Service, error) {
	l := r.log.With(
		zap.Any("operation", "Service.GetService"),
		zap.Any("id", id),
	)

	service, err := r.repo.GetService(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("service not found", zap.Any("error", err))
			return nil, ErrServiceNotFound
		}
		l.Error("get service failed", zap.Any("error", err))
		return nil, err
	}

	return service, nil
}

func (r *Service) GetServices(ctx context.Context) ([]*models.Service, error) {
	l := r.log.With(
		zap.Any("operation", "Service.GetServices"),
	)

	services, err := r.repo.GetServices(ctx)
	if err != nil {
		l.Error("get services failed", zap.Any("error", err))
		return nil, err
	}

	return services, nil
}

func (r *Service) UpdateService(ctx context.Context, service *models.Service) error {
	l := r.log.With(
		zap.Any("operation", "Service.UpdateService"),
		zap.Any("id", service.ID.Hex()),
	)

	if err := r.repo.UpdateService(ctx, service); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("service not found", zap.Any("error", err))
			return ErrServiceNotFound
		}
		l.Error("service update failed", zap.Any("error", err))
		return err
	}

	return nil
}

func (r *Service) DeleteService(ctx context.Context, id string) error {
	l := r.log.With(
		zap.Any("operation", "Service.DeleteService"),
		zap.Any("id", id),
	)

	if err := r.repo.DeleteService(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("service not found", zap.Any("error", err))
			return ErrServiceNotFound
		}
		l.Error("service delete failed", zap.Any("error", err))
		return err
	}

	return nil
}

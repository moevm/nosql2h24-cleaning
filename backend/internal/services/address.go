package services

import (
	"context"
	"errors"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository"
	"go.uber.org/zap"
)

type AddressRepo interface {
	CreateAddress(ctx context.Context, userID string, address *models.Address) (string, error)
	GetAddress(ctx context.Context, userID, addressID string) (*models.AddressWithTimestamp, error)
	GetAddresses(ctx context.Context, userID string) ([]*models.AddressWithTimestamp, error)
	UpdateAddress(ctx context.Context, userID string, address *models.AddressWithTimestamp) error
	DeleteAddress(ctx context.Context, userID string, addressID string) error
}

type AddressService struct {
	log  *zap.Logger
	repo AddressRepo
}

func NewAddressService(
	logger *zap.Logger,
	repo AddressRepo,
) *AddressService {
	return &AddressService{
		log:  logger,
		repo: repo,
	}
}

func (r *AddressService) CreateAddress(ctx context.Context, userID string, address *models.Address) (string, error) {
	l := r.log.With(
		zap.Any("operation", "AddressService.CreateAddress"),
		zap.Any("user_id", userID),
	)

	id, err := r.repo.CreateAddress(ctx, userID, address)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("user not found", zap.Any("error", err))
			return "", ErrUserNotFound
		}
		l.Error("failed to create address", zap.Any("error", err))
		return "", err
	}
	return id, nil
}

func (r *AddressService) GetAddress(ctx context.Context, userID, addressID string) (*models.AddressWithTimestamp, error) {
	l := r.log.With(
		zap.Any("operation", "AddressService.GetAddress"),
		zap.Any("user_id", userID),
		zap.Any("address_id", addressID),
	)

	address, err := r.repo.GetAddress(ctx, userID, addressID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("address not found", zap.Any("error", err))
			return nil, ErrAddressNotFound
		}
		l.Error("get address failed", zap.Any("error", err))
		return nil, err
	}
	return address, nil
}

func (r *AddressService) GetAddresses(ctx context.Context, userID string) ([]*models.AddressWithTimestamp, error) {
	l := r.log.With(
		zap.Any("operation", "AddressService.GetAddresses"),
		zap.Any("user_id", userID),
	)

	addresses, err := r.repo.GetAddresses(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("user not found", zap.Any("error", err))
			return nil, ErrUserNotFound
		}
		l.Error("get addresses failed", zap.Any("error", err))
		return nil, err
	}
	return addresses, nil
}

func (r *AddressService) UpdateAddress(ctx context.Context, userID string, address *models.AddressWithTimestamp) error {
	l := r.log.With(
		zap.Any("operation", "AddressService.UpdateAddress"),
		zap.Any("user_id", userID),
		zap.Any("address_id", address.ID),
	)

	if err := r.repo.UpdateAddress(ctx, userID, address); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("address not found", zap.Any("error", err))
			return ErrAddressNotFound
		}
		l.Error("address update failed", zap.Any("error", err))
		return err
	}
	return nil
}

func (r *AddressService) DeleteAddress(ctx context.Context, userID, addressID string) error {
	l := r.log.With(
		zap.Any("operation", "AddressService.DeleteAddress"),
		zap.Any("user_id", userID),
		zap.Any("address_id", addressID),
	)

	if err := r.repo.DeleteAddress(ctx, userID, addressID); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("address not found", zap.Any("error", err))
			return ErrAddressNotFound
		}
		l.Error("address deletion failed", zap.Any("error", err))
		return err
	}
	return nil
}

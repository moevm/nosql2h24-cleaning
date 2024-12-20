package services

import (
	"context"
	"errors"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"go.uber.org/zap"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUsers(ctx context.Context, filters types.UserFilters) ([]*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
	DeleteWorker(ctx context.Context, id string) error
}

type UserService struct {
	log    *zap.Logger
	hasher Hasher
	repo   UserRepo
}

func NewUserService(
	logger *zap.Logger,
	hasher Hasher,
	repo UserRepo,
) *UserService {
	return &UserService{
		log:    logger,
		hasher: hasher,
		repo:   repo,
	}
}

func (r *UserService) CreateUser(ctx context.Context, user *models.User) (string, error) {
	l := r.log.With(
		zap.Any("operation", "UserService.CreateUser"),
		zap.Any("email", user.Email),
	)

	hash, err := r.hasher.Hash(user.Password)
	if err != nil {
		l.Error("hashing error")
		return "", ErrHashing
	}
	user.Password = hash

	id, err := r.repo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyExist) {
			l.Info(err.Error())
			return "", ErrUserAlreadyExist
		}

		l.Error(
			"failed to create user",
			zap.Any("error", err),
		)
		return "", err
	}
	return id, nil
}

func (r *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	l := r.log.With(
		zap.Any("operation", "UserService.UpdateUser"),
		zap.Any("id", user.ID),
		zap.Any("email", user.Email),
	)

	userFromDB, _ := r.GetUserById(ctx, user.ID.Hex())
	user.PasswordHash = userFromDB.PasswordHash

	if len(user.Password) != 0 {
		hash, err := r.hasher.Hash(user.Password)
		if err != nil {
			l.Error("hashing error")
			return ErrHashing
		}
		user.PasswordHash = hash
	}

	if err := r.repo.UpdateUser(ctx, user); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("user not found", zap.Any("error", err))
			return ErrUserNotFound
		}
		l.Error("update failed", zap.Any("error", err))
		return err
	}
	return nil
}

func (r *UserService) GetUserById(ctx context.Context, id string) (*models.User, error) {
	l := r.log.With(
		zap.Any("operation", "UserService.GetUserById"),
		zap.Any("id", id),
	)

	user, err := r.repo.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidArgument) {

		}
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("user not found", zap.Any("error", err))
			return nil, ErrUserNotFound
		}
		l.Error("get user by id failed", zap.Any("error", err))
		return nil, err
	}
	user.Password = "" // TODO: unset password in repo query
	return user, nil
}

func (r *UserService) GetUsers(ctx context.Context, filters types.UserFilters) ([]*models.User, error) {
	l := r.log.With(
		zap.Any("operation", "UserService.GetUserById"),
	)

	users, err := r.repo.GetUsers(ctx, filters)
	if err != nil {
		l.Error("get users failed", zap.Any("error", err))
		return nil, err
	}
	return users, nil
}

func (r *UserService) DeleteUser(ctx context.Context, id string) error {
	l := r.log.With(
		zap.Any("operation", "UserService.DeleteUser"),
		zap.Any("id", id),
	)

	user, err := r.repo.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("user not found", zap.Any("error", err))
			return ErrUserNotFound
		}
		l.Error("get user by id failed", zap.Any("error", err))
		return err
	}

	if user.UserType == models.UserTypeWorker {
		if err := r.repo.DeleteWorker(ctx, id); err != nil {
			l.Error("delete worker failed", zap.Any("error", err))
			return err
		}
		return nil
	}

	if err := r.repo.DeleteUser(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info("user not found", zap.Any("error", err))
			return ErrUserNotFound
		}
		l.Error("delete user failed", zap.Any("error", err))
		return err
	}
	return nil
}

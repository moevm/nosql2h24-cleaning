package services

import (
	"context"
	"errors"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository"
	"go.uber.org/zap"
)

type JWT interface {
	NewAccessToken(id string) (string, error)
	Parse(token string) (string, error)
}

type Hasher interface {
	Hash(password string) (string, error)
	Compare(password string, hash string) bool
}

type AuthRepo interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
}

type AuthService struct {
	log    *zap.Logger
	jwt    JWT
	hasher Hasher
	repo   AuthRepo
}

func NewAuthService(
	logger *zap.Logger,
	jwt JWT,
	hasher Hasher,
	repo AuthRepo,
) *AuthService {
	return &AuthService{
		log:    logger,
		jwt:    jwt,
		hasher: hasher,
		repo:   repo,
	}
}

func (r *AuthService) Register(ctx context.Context, user *models.User) (string, error) {
	l := r.log.With(
		zap.Any("operation", "AuthService.Register"),
		zap.Any("email", user.Email),
	)

	hash, err := r.hasher.Hash(user.Password)
	if err != nil {
		l.Error("hashing error")
		return "", ErrHashing
	}
	user.PasswordHash = hash

	// On registration user is always a client
	user.UserType = "CLIENT"

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

func (r *AuthService) Login(ctx context.Context, user *models.User) (*models.User, *models.Token, error) {
	l := r.log.With(
		zap.Any("operation", "AuthService.Login"),
		zap.Any("email", user.Email),
	)

	userFromDB, err := r.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info(err.Error())
			return nil, nil, ErrUserNotFound
		}

		l.Error(
			"failed to get user",
			zap.Any("error", err),
		)
		return nil, nil, err
	}

	if !r.hasher.Compare(user.Password, userFromDB.PasswordHash) {
		l.Info("invalid password")
		return nil, nil, ErrIncorrectPassword
	}
	token, err := r.generateJWT(ctx, userFromDB)
	if err != nil {
		l.Error("failed to generate jwt")
		return nil, nil, err
	}
	
	return userFromDB, token, nil
}

func (r *AuthService) UserExist(ctx context.Context, id string) error {
	l := r.log.With(
		zap.Any("operation", "AuthService.UserExist"),
		zap.Any("id", id),
	)

	if _, err := r.repo.GetUserById(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			l.Info(err.Error())
			return ErrUserNotFound
		}

		l.Error(
			"failed to get user",
			zap.Any("error", err),
		)
		return err
	}
	return nil
}

func (r *AuthService) Logout(ctx context.Context, refreshToken string) error {
	// TODO: Implement me
	return nil
}

func (r *AuthService) generateJWT(_ context.Context, user *models.User) (*models.Token, error) {
	l := r.log.With(
		zap.Any("operation", "AuthService.Login"),
		zap.Any("email", user.Email),
	)

	access, err := r.jwt.NewAccessToken(user.ID.Hex())
	if err != nil {
		l.Info("failed to sign token")
		return nil, err
	}
	return &models.Token{
		Access: access,
	}, nil
}

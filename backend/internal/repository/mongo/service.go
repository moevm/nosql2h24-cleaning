package mongorepo

import (
	"context"
	"errors"
	"time"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ServiceRepo struct {
	collection *mongo.Collection
}

func NewServiceRepo(db *mongo.Database) *ServiceRepo {
	return &ServiceRepo{
		collection: db.Collection(ServicesCollection),
	}
}

func (r *ServiceRepo) CreateService(ctx context.Context, service *models.Service) (string, error) {
	service.CreatedAt = time.Now()

	res, err := r.collection.InsertOne(ctx, service)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", repository.ErrAlreadyExist
		}
		return "", err
	}

	return res.InsertedID.(bson.ObjectID).Hex(), nil
}

func (r *ServiceRepo) GetService(ctx context.Context, id string) (*models.Service, error) {
	_id, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, repository.ErrInvalidArgument
	}

	filter := bson.D{{Key: "_id", Value: _id}}

	var service models.Service
	if err := r.collection.FindOne(ctx, filter).Decode(&service); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &service, nil
}

func (r *ServiceRepo) UpdateService(ctx context.Context, service *models.Service) error {
	update := bson.D{
		{Key: "$set", Value: service},
		{Key: "$currentDate", Value: bson.D{
			{Key: "updated_at", Value: true},
		}},
	}

	res, err := r.collection.UpdateByID(ctx, service.ID, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return repository.ErrNotFound
		}
		return err
	}
	if res.MatchedCount == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (r *ServiceRepo) DeleteService(ctx context.Context, id string) error {
	_id, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return repository.ErrInvalidArgument
	}

	filter := bson.D{{Key: "_id", Value: _id}}

	res, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return repository.ErrNotFound
		}
		return err
	}
	if res.DeletedCount == 0 {
		return repository.ErrNotFound
	}

	return nil
}

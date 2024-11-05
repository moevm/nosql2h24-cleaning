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

type OrderRepo struct {
	collection *mongo.Collection
}

func NewOrderRepo(db *mongo.Database) *OrderRepo {
	return &OrderRepo{
		collection: db.Collection(OrdersCollection),
	}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, order *models.Order) (string, error) {
	order.CreatedAt = time.Now()

	res, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", repository.ErrAlreadyExist
		}
		return "", err
	}

	return res.InsertedID.(bson.ObjectID).Hex(), nil
}

func (r *OrderRepo) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	_id, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, repository.ErrInvalidArgument
	}

	filter := bson.D{{Key: "_id", Value: _id}}

	var order models.Order
	if err := r.collection.FindOne(ctx, filter).Decode(&order); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepo) UpdateOrder(ctx context.Context, order *models.Order) error {
	update := bson.D{
		{Key: "$set", Value: order},
		{Key: "$currentDate", Value: bson.D{
			{Key: "updated_at", Value: true},
		}},
	}

	res, err := r.collection.UpdateByID(ctx, order.ID, update)
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

func (r *OrderRepo) DeleteOrder(ctx context.Context, id string) error {
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

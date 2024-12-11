package mongorepo

import (
	"context"
	"errors"
	"time"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
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

func (r *OrderRepo) ImportOrders(ctx context.Context, orders []models.Order) error {
	count, err := r.collection.EstimatedDocumentCount(ctx)
	if err != nil {
		return err
	}
	if count != 0 {
		return repository.ErrAlreadyExist
	}
	_, err = r.collection.InsertMany(ctx, orders)
	return err
}

func (r *OrderRepo) ExportOrders(ctx context.Context) ([]models.Order, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	orders := make([]models.Order, 0)
	for cursor.Next(ctx) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepo) Drop(ctx context.Context) error {
	return r.collection.Drop(ctx)
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

func (r *OrderRepo) GetOrderById(ctx context.Context, id string) (*models.Order, error) {
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

func (r *OrderRepo) GetUserOrders(ctx context.Context, userID string) ([]*models.Order, error) {
	_id, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, repository.ErrInvalidArgument
	}

	filter := bson.D{{Key: "user_id", Value: _id}}

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var orders []*models.Order
	if err := cur.All(ctx, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepo) GetOrders(ctx context.Context, query types.SearchParams) ([]*models.Order, error) {
	filter := bson.D{}

	// Not sure if this is the best way to handle this
	if query.UserID != "" {
		_id, err := bson.ObjectIDFromHex(query.UserID)
		if err != nil {
			return nil, repository.ErrInvalidArgument
		}
		filter = append(filter, bson.E{Key: "user_id", Value: _id})
	}
	if len(query.WorkersID) != 0 {
		_ids := make([]bson.ObjectID, 0, len(query.WorkersID))
		for _, id := range query.WorkersID {
			_id, err := bson.ObjectIDFromHex(id)
			if err != nil {
				return nil, repository.ErrInvalidArgument
			}
			_ids = append(_ids, _id)
		}
		filter = append(filter, bson.E{Key: "workers", Value: bson.D{{Key: "$in", Value: _ids}}})
	}
	if len(query.Statuses) != 0 {
		filter = append(filter, bson.E{Key: "status", Value: bson.D{{Key: "$in", Value: query.Statuses}}})
	}

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var orders []*models.Order
	if err := cur.All(ctx, &orders); err != nil {
		return nil, err
	}

	return orders, nil
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

package mongorepo

import (
	"context"
	"errors"
	"time"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/mongo/search"
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

func (r *OrderRepo) GetOrders(ctx context.Context, query types.OrderFilters) ([]*models.Order, error) {
	filter := search.Filter{}
	// user info
	filter.AddEqualObjectID("user_id", query.UserID)

	filter.AddRegex("user.name", query.User.Name)
	filter.AddRegex("user.surname", query.User.Surname)
	filter.AddRegex("user.patronymic", query.User.Patronymic)
	filter.AddRegex("user.email", query.User.Email)

	// workers
	filter.AddInObjectIDs("workers", query.WorkersID)

	// address
	filter.AddRegex("address.city", query.Address.City)
	filter.AddRegex("address.street", query.Address.Street)
	filter.AddRegex("address.building", query.Address.Building)
	filter.AddRegex("address.entrance", query.Address.Entrance)
	filter.AddRegex("address.floor", query.Address.Floor)
	filter.AddRegex("address.door_number", query.Address.DoorNumber)

	filter.AddIn("status", query.Statuses)
	filter.ContainsAll("services", query.Services)

	// time
	filter.AddTimeIterval("date_time", query.DateTime)

	cur, err := r.collection.Aggregate(ctx, bson.A{
		bson.D{{
			Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "user_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "user"},
			},
		}},
		bson.D{{
			Key:   "$match",
			Value: filter.ToBson(),
		}},
	})
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
	order.UpdatedAt = time.Now()
	update := bson.D{
		{Key: "$set", Value: order},
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

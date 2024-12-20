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
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/sync/errgroup"
)

type UserRepo struct {
	collection *mongo.Collection
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	collection := db.Collection(UsersCollection)
	// create unique constraint
	collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	return &UserRepo{
		collection: collection,
	}
}

func (r *UserRepo) ImportUsers(ctx context.Context, users []models.User) error {
	count, err := r.collection.EstimatedDocumentCount(ctx)
	if err != nil {
		return err
	}
	if count != 0 {
		return repository.ErrAlreadyExist
	}
	_, err = r.collection.InsertMany(ctx, users)
	return err
}

func (r *UserRepo) ExportUsers(ctx context.Context) ([]models.User, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	users := make([]models.User, 0)
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepo) Drop(ctx context.Context) error {
	return r.collection.Drop(ctx)
}

func (r *UserRepo) CreateUser(ctx context.Context, user *models.User) (string, error) {
	user.CreatedAt = time.Now()

	res, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", repository.ErrAlreadyExist
		}
		return "", err
	}

	return res.InsertedID.(bson.ObjectID).Hex(), nil
}

func (r *UserRepo) GetUserById(ctx context.Context, id string) (*models.User, error) {
	_id, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, repository.ErrInvalidArgument
	}

	filter := bson.D{{Key: "_id", Value: _id}}

	var user models.User
	if err := r.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var user models.User
	if err := r.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUsers(ctx context.Context, filters types.UserFilters) ([]*models.User, error) {
	filter := search.Filter{}
	filter.AddEqual("user_type", filters.UserType)
	filter.AddRegex("name", filters.Name)
	filter.AddRegex("surname", filters.Surname)
	filter.AddRegex("email", filters.Email)
	filter.AddRegex("phone_number", filters.PhoneNumber)
	filter.AddTimeIterval("created_at", filters.CreatedAt)

	// worker filters
	filter.AddNumberRange("orders_count", filters.OrdersCount)

	cursor, err := r.collection.Find(ctx, filter.ToBson())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	users := make([]*models.User, 0)
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, user *models.User) error {
	update := bson.D{
		{Key: "$set", Value: user},
		{Key: "$currentDate", Value: bson.D{
			{Key: "updated_at", Value: true},
		}},
	}

	res, err := r.collection.UpdateByID(
		ctx,
		user.ID,
		update,
	)
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

func (r *UserRepo) DeleteUser(ctx context.Context, id string) error {
	_id, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return repository.ErrInvalidArgument
	}

	filter := bson.D{{Key: "_id", Value: _id}}

	res, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (r *UserRepo) DeleteWorker(ctx context.Context, id string) error {
	_id, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return repository.ErrInvalidArgument
	}

	eg := &errgroup.Group{}
	// gorutine 1: delete from user
	eg.Go(func() error {
		return r.DeleteUser(ctx, id)
	})

	// gorutine 2: delete from orders where status != models.OrderStatusDone
	eg.Go(func() error {
		filter := search.Filter{}
		filter.AddNotEqual("status", models.OrderStatusDone)
		filter.AddInObjectIDs("workers", []string{id})

		update := bson.A{
			bson.M{"$set": bson.M{
				"workers": bson.M{
					"$filter": bson.M{
						"input": "$workers", // Исходный массив workers
						"as":    "worker",   // Псевдоним для элементов массива
						"cond": bson.M{
							"$ne": bson.A{"$$worker", _id}, // Убираем _id из массива
						},
					},
				},
				"status_logs": bson.M{
					"$cond": bson.M{
						"if": bson.M{
							"$eq": bson.A{
								"$status", models.OrderStatusAccepted, // Проверка, что статус равен "active"
							},
						},
						"then": bson.M{
							"$concatArrays": bson.A{
								"$status_logs", // Текущий массив status_log
								bson.A{
									// Новый элемент, который будет добавлен
									bson.M{
										"new_status":  models.OrderStatusNew,
										"prev_status": "$status",
										"created_at":  time.Now(),
									},
								},
							},
						},
						"else": "$status_logs", // Если условие не выполнено, оставляем текущий status_log
					},
				},
				"status": bson.M{
					"$cond": bson.M{
						"if": bson.M{
							"$eq": bson.A{
								"$status", models.OrderStatusAccepted, // Проверка, что статус равен "active"
							},
						},
						"then": models.OrderStatusNew, // Если количество работников меньше требуемого, обновляем статус
						"else": "$status",             // В противном случае оставляем текущий статус
					},
				},
				"updated_at": time.Now(),
			},
			},
		}

		_, err := r.collection.Database().Collection(OrdersCollection).UpdateMany(ctx, filter.ToBson(), update)
		return err
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

// Operations with user addresses
func (r *UserRepo) CreateAddress(ctx context.Context, userID string, address *models.Address) (string, error) {
	_id, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return "", repository.ErrInvalidArgument
	}
	address.ID = bson.NewObjectID()

	withTimestamp := models.AddressWithTimestamp{}
	withTimestamp.Address = *address
	withTimestamp.CreatedAt = time.Now()

	update := bson.D{
		bson.E{
			Key: "$push",
			Value: bson.D{
				{Key: "addresses", Value: withTimestamp},
			},
		},
		bson.E{
			Key: "$currentDate",
			Value: bson.D{
				{Key: "updated_at", Value: true},
			},
		},
	}

	res, err := r.collection.UpdateByID(ctx, _id, update, options.Update().SetUpsert(true))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", repository.ErrNotFound
		}
		return "", err
	}
	if res.ModifiedCount == 0 {
		return "", repository.ErrNotFound
	}
	return address.ID.Hex(), nil
}

func (r *UserRepo) GetAddress(ctx context.Context, userID, addressID string) (*models.AddressWithTimestamp, error) {
	user_id, userErr := bson.ObjectIDFromHex(userID)
	address_id, addressErr := bson.ObjectIDFromHex(addressID)
	if userErr != nil || addressErr != nil {
		return nil, repository.ErrInvalidArgument
	}

	filter := bson.D{
		{Key: "_id", Value: user_id},
		{Key: "addresses._id", Value: address_id},
	}
	projection := bson.D{
		{Key: "addresses", Value: bson.D{
			{Key: "$elemMatch", Value: bson.D{
				{Key: "_id", Value: address_id},
			}},
		}},
	}

	var user models.User
	if err := r.collection.FindOne(
		ctx,
		filter,
		options.FindOne().SetProjection(projection),
	).Decode(&user); err != nil {
		return nil, err
	}

	if len(user.Addresses) == 0 {
		return nil, repository.ErrNotFound
	}
	return user.Addresses[0], nil
}

func (r *UserRepo) GetAddresses(ctx context.Context, userID string) ([]*models.AddressWithTimestamp, error) {
	_id, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, repository.ErrInvalidArgument
	}

	filter := bson.D{
		{Key: "_id", Value: _id},
	}
	projection := bson.D{
		{Key: "addresses", Value: 1},
	}

	var user models.User
	if err := r.collection.FindOne(
		ctx,
		filter,
		options.FindOne().SetProjection(projection),
	).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return user.Addresses, nil
}

func (r *UserRepo) UpdateAddress(ctx context.Context, userID string, address *models.AddressWithTimestamp) error {
	_id, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return repository.ErrInvalidArgument
	}
	address.UpdatedAt = time.Now()

	filter := bson.D{
		{Key: "_id", Value: _id},
		{Key: "addresses._id", Value: address.ID},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "addresses.$", Value: address},
		}},
	}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return repository.ErrNotFound
		}
		return err
	}
	if res.ModifiedCount == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (r *UserRepo) DeleteAddress(ctx context.Context, userID string, addressID string) error {
	user_id, userErr := bson.ObjectIDFromHex(userID)
	address_id, addressErr := bson.ObjectIDFromHex(addressID)
	if userErr != nil || addressErr != nil {
		return repository.ErrInvalidArgument
	}

	filter := bson.D{
		{Key: "_id", Value: user_id},
	}
	update := bson.D{
		{Key: "$pull", Value: bson.D{
			{Key: "addresses", Value: bson.D{
				{Key: "_id", Value: address_id},
			}},
		}},
	}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return repository.ErrNotFound
		}
		return err
	}
	if res.ModifiedCount == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *UserRepo) IncrementWorkersOrderCounts(ctx context.Context, workersIDs ...string) error {
	filter := search.Filter{}
	filter.AddInObjectIDs("_id", workersIDs)

	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "orders_count", Value: 1},
		}},
	}

	res, err := r.collection.UpdateMany(ctx, filter.ToBson(), update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return repository.ErrNotFound
		}
		return err
	}
	if res.ModifiedCount == 0 {
		return repository.ErrNotFound
	}
	return nil
}

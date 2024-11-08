package mongorepo

import (
	"context"
	"errors"
	"time"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserRepo struct {
	collection *mongo.Collection
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{
		collection: db.Collection(UsersCollection),
	}
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

// TODO: third argument query
func (r *UserRepo) GetUsers(ctx context.Context, userType string) ([]*models.User, error) {
	filter := bson.D{{Key: "user_type", Value: userType}}

	cursor, err := r.collection.Find(ctx, filter)
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

// Operations with user addresses
func (r *UserRepo) CreateAddress(ctx context.Context, userID string, address *models.Address) (string, error) {
	_id, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return "", repository.ErrInvalidArgument
	}
	address.ID = bson.NewObjectID()
	address.CreatedAt = time.Now()

	update := bson.D{
		bson.E{
			Key: "$push",
			Value: bson.D{
				{Key: "addresses", Value: address},
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

func (r *UserRepo) GetAddress(ctx context.Context, userID, addressID string) (*models.Address, error) {
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

func (r *UserRepo) GetAddresses(ctx context.Context, userID string) ([]*models.Address, error) {
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

func (r *UserRepo) UpdateAddress(ctx context.Context, userID string, address *models.Address) error {
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

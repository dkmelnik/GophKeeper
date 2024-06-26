package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
	"github.com/dkmelnik/GophKeeper/internal/domain/errs"
)

// UserRepository is a repository for handling user data in MongoDB.
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db.Collection("users")}
}

// Save inserts a new user entity into the MongoDB collection.
// It returns the updated entity.User entity or an error if insertion fails.
func (r *UserRepository) Save(ctx context.Context, user entity.User) (entity.User, error) {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("%w: can't save user: %v", err, user)
	}
	insertedID := result.InsertedID.(primitive.ObjectID)
	user.ID = entity.FromObjectID(insertedID)
	return user, nil
}

// IsEntryByLogin check exist entity in db by login.
// It returns bool or an error if retrieval fails.
func (r *UserRepository) IsEntryByLogin(ctx context.Context, login string) (bool, error) {
	filter := bson.D{{"login", login}}
	err := r.collection.FindOne(ctx, filter).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// FindOneByLogin retrieves text entity associated with a specific login.
// It returns entity.User entity matching the login or an error if retrieval fails.
func (r *UserRepository) FindOneByLogin(ctx context.Context, login string) (entity.User, error) {
	var result entity.User
	filter := bson.D{{"login", login}}
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return entity.User{}, fmt.Errorf("%w: can't find user: %q", errs.ErrNotFound, login)
		}
		return entity.User{}, err
	}

	return result, nil
}

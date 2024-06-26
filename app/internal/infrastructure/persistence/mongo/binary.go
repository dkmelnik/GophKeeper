package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
)

// BinaryRepository is a repository for handling binary data in MongoDB.
type BinaryRepository struct {
	collection *mongo.Collection
}

// NewBinaryRepository creates a new BinaryRepository instance.
func NewBinaryRepository(db *mongo.Database) *BinaryRepository {
	return &BinaryRepository{db.Collection("binary")}
}

// Save inserts a new binary entity into the MongoDB collection.
// It returns the updated entity.Binary entity or an error if insertion fails.
func (r *BinaryRepository) Save(ctx context.Context, binary entity.Binary) (entity.Binary, error) {
	result, err := r.collection.InsertOne(ctx, binary)
	if err != nil {
		return entity.Binary{}, fmt.Errorf("%w: can't save binary", err)
	}
	insertedID := result.InsertedID.(primitive.ObjectID)
	binary.ID = entity.FromObjectID(insertedID)
	return binary, nil
}

// FindByUserID retrieves binary entities associated with a specific user ID.
// It returns a slice of entity.Binary entities matching the user ID or an error if retrieval fails.
func (r *BinaryRepository) FindByUserID(ctx context.Context, userID entity.ID) ([]entity.Binary, error) {
	oid, err := userID.ObjectID()
	if err != nil {
		return nil, fmt.Errorf("%w: can't convert user_id: %s", err, userID)
	}
	filter := bson.M{"user_id": oid}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%w: can't find binary by user_id: %s", err, userID)
	}
	defer cursor.Close(ctx)

	var binary []entity.Binary
	for cursor.Next(ctx) {
		var b entity.Binary
		if err = cursor.Decode(&b); err != nil {
			return nil, fmt.Errorf("can't decode text: %w", err)
		}
		binary = append(binary, b)
	}

	if err = cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return binary, nil
}

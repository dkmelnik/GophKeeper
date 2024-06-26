package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
)

// TextRepository is a repository for handling text data in MongoDB.
type TextRepository struct {
	collection *mongo.Collection
}

// NewTextRepository creates a new TextRepository instance.
func NewTextRepository(db *mongo.Database) *TextRepository {
	return &TextRepository{db.Collection("texts")}
}

// Save inserts a new text entity into the MongoDB collection.
// It returns the updated entity.Text entity or an error if insertion fails.
func (r *TextRepository) Save(ctx context.Context, text entity.Text) (entity.Text, error) {
	result, err := r.collection.InsertOne(ctx, text)
	if err != nil {
		return entity.Text{}, fmt.Errorf("%w: can't save text: %v", err, text)
	}
	insertedID := result.InsertedID.(primitive.ObjectID)
	text.ID = entity.FromObjectID(insertedID)
	return text, nil
}

// FindByUserID retrieves text entities associated with a specific user ID.
// It returns a slice of entity.Text entities matching the user ID or an error if retrieval fails.
func (r *TextRepository) FindByUserID(ctx context.Context, userID entity.ID) ([]entity.Text, error) {
	oid, err := userID.ObjectID()
	if err != nil {
		return nil, fmt.Errorf("%w: can't convert user_id: %s", err, userID)
	}
	filter := bson.M{"user_id": oid}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%w: can't find texts by user_id: %s", err, userID)
	}
	defer cursor.Close(ctx)

	var texts []entity.Text
	for cursor.Next(ctx) {
		var text entity.Text
		if err = cursor.Decode(&text); err != nil {
			return nil, fmt.Errorf("can't decode text: %w", err)
		}
		texts = append(texts, text)
	}

	if err = cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return texts, nil
}

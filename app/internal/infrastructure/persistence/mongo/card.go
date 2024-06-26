package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
)

// CardRepository is a repository for handling card data in MongoDB.
type CardRepository struct {
	collection *mongo.Collection
}

// NewCardRepository creates a new CardRepository instance.
func NewCardRepository(db *mongo.Database) *CardRepository {
	return &CardRepository{db.Collection("cards")}
}

// Save inserts a new card entity into the MongoDB collection.
// It returns the updated entity.Card entity or an error if insertion fails.
func (r *CardRepository) Save(ctx context.Context, card entity.Card) (entity.Card, error) {
	result, err := r.collection.InsertOne(ctx, card)
	if err != nil {
		return entity.Card{}, fmt.Errorf("%w: can't save card: %v", err, card)
	}
	insertedID := result.InsertedID.(primitive.ObjectID)
	card.ID = entity.FromObjectID(insertedID)
	return card, nil
}

// FindByUserID retrieves card entities associated with a specific user ID.
// It returns a slice of entity.Card entities matching the user ID or an error if retrieval fails.
func (r *CardRepository) FindByUserID(ctx context.Context, userID entity.ID) ([]entity.Card, error) {
	oid, err := userID.ObjectID()
	if err != nil {
		return nil, fmt.Errorf("%w: can't convert user_id: %s", err, userID)
	}
	filter := bson.M{"user_id": oid}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%w: can't find cards by user_id: %s", err, userID)
	}
	defer cursor.Close(ctx)

	var cards []entity.Card
	for cursor.Next(ctx) {
		var card entity.Card
		if err = cursor.Decode(&card); err != nil {
			return nil, fmt.Errorf("can't decode card: %w", err)
		}
		cards = append(cards, card)
	}

	if err = cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return cards, nil
}

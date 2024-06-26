package entity

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ID represents an identifier that can be marshaled to and from MongoDB ObjectID.
type ID string

// FromObjectID converts a primitive.ObjectID to ID.
func FromObjectID(o primitive.ObjectID) ID {
	return ID(o.Hex())
}

// ObjectID converts ID to primitive.ObjectID.
func (id *ID) ObjectID() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(string(*id))
}

// String returns the string representation of ID.
func (id *ID) String() string {
	return string(*id)
}

// MarshalBSONValue implements the bson.Marshaler interface for Email
func (e ID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	o, err := primitive.ObjectIDFromHex(e.String())
	if err != nil {
		return 0, nil, err
	}
	return bson.MarshalValue(o)
}

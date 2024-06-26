package entity

// Binary represents binary data stored in MongoDB.
type Binary struct {
	ID            ID                `bson:"_id,omitempty"`  // Unique identifier for the binary content
	UserID        ID                `bson:"user_id"`        // ID of the user associated with the binary content
	BinaryContent string            `bson:"binary_content"` // Actual binary content stored as a string (base64 or similar)
	Metadata      map[string]string `bson:"metadata"`       // Additional metadata associated with the binary content
}

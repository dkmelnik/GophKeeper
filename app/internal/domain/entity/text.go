package entity

// Text represents textual data stored in MongoDB.
type Text struct {
	ID          ID                `bson:"_id,omitempty"` // Unique identifier for the text content
	UserID      ID                `bson:"user_id"`       // ID of the user associated with the text content
	TextContent string            `bson:"text_content"`  // The actual textual content
	Metadata    map[string]string `bson:"metadata"`      // Additional metadata associated with the text content
}

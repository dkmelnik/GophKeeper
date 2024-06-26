package entity

// User represents user information stored in MongoDB.
type User struct {
	ID       ID     `bson:"_id,omitempty"` // Unique identifier for the user
	Login    string `bson:"login"`         // User's login username
	Password string `bson:"password"`      // User's password (hashed or encrypted)
}

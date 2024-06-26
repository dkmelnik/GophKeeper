package entity

// Card represents credit card information stored in MongoDB.
type Card struct {
	ID         ID                `bson:"_id,omitempty"` // Unique identifier for the card
	UserID     ID                `bson:"user_id"`       // ID of the user associated with the card
	CardNumber string            `bson:"card_number"`   // Card number (masked or full)
	ExpiryDate string            `bson:"expiry_date"`   // Expiration date of the card (e.g., "MM/YYYY")
	CVV        string            `bson:"cvv"`           // Card Verification Value (CVV) security code
	Metadata   map[string]string `bson:"metadata"`      // Additional metadata associated with the card
}

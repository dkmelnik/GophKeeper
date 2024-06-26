package dto

import (
	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/utils"
)

// CardDetails represents specific details extracted from entity.Card.
type CardDetails struct {
	CardNumber string // CardNumber holds the card number information.
	ExpiryDate string // ExpiryDate holds the card expiry date.
	CVV        string // CVV holds the card CVV information with masking.
	Metadata   string // Metadata holds the metadata converted to a string format.
}

// ToCardDetails converts an entity.Card instance to CardDetails.
func ToCardDetails(card entity.Card) CardDetails {
	return CardDetails{
		CardNumber: card.CardNumber,
		ExpiryDate: card.ExpiryDate,
		CVV:        maskCVV(card.CVV),                // Mask CVV for security
		Metadata:   utils.MapToString(card.Metadata), // Convert metadata map to string format
	}
}

// maskCVV masks the CVV by replacing the middle character with "*".
func maskCVV(cvv string) string {
	if len(cvv) != 3 {
		return cvv // Return original CVV if not exactly 3 characters long
	}
	return string(cvv[0]) + "*" + string(cvv[2]) // Mask CVV
}

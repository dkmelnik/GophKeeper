package dto

import (
	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/utils"
)

// TextDetails represents specific details extracted from entity.Text.
type TextDetails struct {
	TextContent string // TextContent holds the text content data.
	Metadata    string // Metadata holds the metadata converted to a string format.
}

// ToTextDetails converts an entity.Text instance to TextDetails.
func ToTextDetails(text entity.Text) TextDetails {
	return TextDetails{
		TextContent: text.TextContent,
		Metadata:    utils.MapToString(text.Metadata), // Convert metadata map to string format
	}
}

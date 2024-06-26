package dto

import (
	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/utils"
)

// BinaryDetails represents specific details extracted from entity.Binary.
type BinaryDetails struct {
	BinaryContent string // BinaryContent holds the binary content data.
	Metadata      string // Metadata holds the metadata converted to a string format.
}

// ToBinaryDetails converts an entity.Binary instance to BinaryDetails.
func ToBinaryDetails(binary entity.Binary) BinaryDetails {
	return BinaryDetails{
		BinaryContent: binary.BinaryContent,
		Metadata:      utils.MapToString(binary.Metadata), // Convert metadata map to string format
	}
}

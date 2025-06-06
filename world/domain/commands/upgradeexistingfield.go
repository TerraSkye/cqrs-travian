package commands

import (
	"github.com/google/uuid"
	"time"
)

type UpgradeExistingFieldCommand struct {
	WorldID   uuid.UUID
	VillageID int64
	FieldID   int64

	// cost
	Wood int64
	Clay int64
	Iron int64
	Crop int64
	//
	CropConsumption int64
	CulturePoint    int64
	//TTL
	Duration time.Duration
}

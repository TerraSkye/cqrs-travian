package events

import (
	"github.com/google/uuid"
)

var _ infra.Event[uuid.UUID] = (*FieldUpgradeFinishedEvent)(nil)

type FieldUpgradeFinishedEvent struct {
	WorldID         uuid.UUID
	VillageID       uuid.UUID
	UpgradeID       uuid.UUID
	TileID          int64
	CropConsumption int64
	CulturePoint    int64
}

func (e *FieldUpgradeFinishedEvent) AggregateID() uuid.UUID {
	return e.VillageID
}

package events

import (
	"github.com/google/uuid"
)

var _ infra.Event[uuid.UUID] = (*BuildingUpgradeStartedEvent)(nil)

type BuildingUpgradeStartedEvent struct {
	WorldID    uuid.UUID
	VillageID  uuid.UUID
	BuildingID int64
	//ticks until completion
	Duration int64
}

func (e *BuildingUpgradeStartedEvent) AggregateID() uuid.UUID {
	return e.VillageID
}

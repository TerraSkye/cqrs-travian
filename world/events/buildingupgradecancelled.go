package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*BuildingUpgradeCancelledEvent)(nil)

type BuildingUpgradeCancelledEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	UpgradeID uuid.UUID
}

func (e *BuildingUpgradeCancelledEvent) AggregateID() uuid.UUID {
	return e.VillageID
}

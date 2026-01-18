package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*BuildingQueuedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&BuildingQueuedEvent{})
}

type BuildingQueuedEvent struct {
	WorldID       uuid.UUID
	VillageID     uuid.UUID
	QueueID       uuid.UUID
	BuildingType  support.BuildingType
	Slot          int
	FromLevel     int
	ToLevel       int
	Cost          support.Resources
	DurationTicks int64 // Duration in game ticks
}

func (e BuildingQueuedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e BuildingQueuedEvent) AggregateID() string {
	return e.VillageID.String()
}

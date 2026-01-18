package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*BuildingCompletedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&BuildingCompletedEvent{})
}

type BuildingCompletedEvent struct {
	WorldID      uuid.UUID
	VillageID    uuid.UUID
	QueueID      uuid.UUID
	BuildingType support.BuildingType
	Slot         int
	NewLevel     int
}

func (e BuildingCompletedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e BuildingCompletedEvent) AggregateID() string {
	return e.VillageID.String()
}

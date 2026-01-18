package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*BuildingStartedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&BuildingStartedEvent{})
}

type BuildingStartedEvent struct {
	WorldID      uuid.UUID
	VillageID    uuid.UUID
	QueueID      uuid.UUID
	BuildingType support.BuildingType
	Slot         int
	FromLevel    int
	ToLevel      int
	StartTick    int64 // Tick when construction started
	EndTick      int64 // Tick when construction completes
}

func (e BuildingStartedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e BuildingStartedEvent) AggregateID() string {
	return e.VillageID.String()
}

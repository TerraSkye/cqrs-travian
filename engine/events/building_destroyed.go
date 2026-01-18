package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*BuildingDestroyedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&BuildingDestroyedEvent{})
}

type BuildingDestroyedEvent struct {
	WorldID       uuid.UUID
	VillageID     uuid.UUID
	AttackerID    uuid.UUID
	BuildingType  support.BuildingType
	Slot          int
	FromLevel     int
	ToLevel       int
	CatapultsUsed int
}

func (e BuildingDestroyedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e BuildingDestroyedEvent) AggregateID() string {
	return e.VillageID.String()
}

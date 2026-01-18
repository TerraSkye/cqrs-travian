package events

import (
	support2 "cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*BuildingCancelledEvent)(nil)

func init() {
	cqrs.RegisterEvent(&BuildingCancelledEvent{})
}

type BuildingCancelledEvent struct {
	WorldID           uuid.UUID
	VillageID         uuid.UUID
	QueueID           uuid.UUID
	BuildingType      support2.BuildingType
	Slot              int
	RefundedResources support2.Resources
}

func (e BuildingCancelledEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e BuildingCancelledEvent) AggregateID() string {
	return e.VillageID.String()
}

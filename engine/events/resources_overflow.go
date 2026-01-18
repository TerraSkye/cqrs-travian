package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ResourcesOverflowEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ResourcesOverflowEvent{})
}

type ResourcesOverflowEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	Lost      support.Resources
	Capacity  support.Resources
}

func (e ResourcesOverflowEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ResourcesOverflowEvent) AggregateID() string {
	return e.VillageID.String()
}

package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ResourcesProducedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ResourcesProducedEvent{})
}

type ResourcesProducedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	Amount    support.Resources
	TickCount int64
}

func (e ResourcesProducedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ResourcesProducedEvent) AggregateID() string {
	return e.VillageID.String()
}

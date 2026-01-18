package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ResourcesConsumedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ResourcesConsumedEvent{})
}

type ResourcesConsumedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	Amount    support.Resources
	Reason    string
}

func (e ResourcesConsumedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ResourcesConsumedEvent) AggregateID() string {
	return e.VillageID.String()
}

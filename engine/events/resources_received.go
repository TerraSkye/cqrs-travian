package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ResourcesReceivedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ResourcesReceivedEvent{})
}

type ResourcesReceivedEvent struct {
	WorldID       uuid.UUID
	VillageID     uuid.UUID
	FromVillageID uuid.UUID
	Amount        support.Resources
	Source        string
}

func (e ResourcesReceivedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ResourcesReceivedEvent) AggregateID() string {
	return e.VillageID.String()
}

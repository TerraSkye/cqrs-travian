package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*NPCTradeCompletedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&NPCTradeCompletedEvent{})
}

type NPCTradeCompletedEvent struct {
	WorldID      uuid.UUID
	VillageID    uuid.UUID
	OldResources support.Resources
	NewResources support.Resources
	GoldSpent    int
}

func (e NPCTradeCompletedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e NPCTradeCompletedEvent) AggregateID() string {
	return e.VillageID.String()
}

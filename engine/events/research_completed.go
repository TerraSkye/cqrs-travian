package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ResearchCompletedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ResearchCompletedEvent{})
}

type ResearchCompletedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	TroopType support.TroopType
}

func (e ResearchCompletedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ResearchCompletedEvent) AggregateID() string {
	return e.VillageID.String()
}

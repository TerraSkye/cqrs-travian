package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ResearchStartedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ResearchStartedEvent{})
}

type ResearchStartedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	TroopType support.TroopType
	Cost      support.Resources
	StartTick int64 // Tick when research started
	EndTick   int64 // Tick when research completes
}

func (e ResearchStartedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ResearchStartedEvent) AggregateID() string {
	return e.VillageID.String()
}

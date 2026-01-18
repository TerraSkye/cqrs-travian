package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ResourcesRaidedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ResourcesRaidedEvent{})
}

type ResourcesRaidedEvent struct {
	WorldID         uuid.UUID
	AttackerVillage uuid.UUID
	DefenderVillage uuid.UUID
	Amount          support.Resources
	CarryCapacity   int64
}

func (e ResourcesRaidedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ResourcesRaidedEvent) AggregateID() string {
	return e.DefenderVillage.String()
}

package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ReinforcementRecalledEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ReinforcementRecalledEvent{})
}

type ReinforcementRecalledEvent struct {
	WorldID       uuid.UUID
	MovementID    uuid.UUID
	FromVillageID uuid.UUID
	ToVillageID   uuid.UUID
	Troops        support.Troops
	ReturnTick    int64 // Tick when troops will return home
}

func (e ReinforcementRecalledEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ReinforcementRecalledEvent) AggregateID() string {
	return e.ToVillageID.String()
}

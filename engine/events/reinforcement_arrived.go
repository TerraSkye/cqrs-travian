package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ReinforcementArrivedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ReinforcementArrivedEvent{})
}

type ReinforcementArrivedEvent struct {
	WorldID       uuid.UUID
	MovementID    uuid.UUID
	FromVillageID uuid.UUID
	FromPlayerID  uuid.UUID
	ToVillageID   uuid.UUID
	Troops        support.Troops
}

func (e ReinforcementArrivedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ReinforcementArrivedEvent) AggregateID() string {
	return e.ToVillageID.String()
}

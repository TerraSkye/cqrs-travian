package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TroopsArrivedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TroopsArrivedEvent{})
}

type TroopsArrivedEvent struct {
	WorldID       uuid.UUID
	MovementID    uuid.UUID
	FromVillageID uuid.UUID
	ToVillageID   uuid.UUID
	Troops        support.Troops
	MovementType  support.MovementType
}

func (e TroopsArrivedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TroopsArrivedEvent) AggregateID() string {
	return e.ToVillageID.String()
}

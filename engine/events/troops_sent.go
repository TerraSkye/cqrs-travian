package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TroopsSentEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TroopsSentEvent{})
}

type TroopsSentEvent struct {
	WorldID       uuid.UUID
	MovementID    uuid.UUID
	FromVillageID uuid.UUID
	ToVillageID   uuid.UUID
	ToX           int
	ToY           int
	Troops        support.Troops
	MovementType  support.MovementType
	DepartureTick int64 // Tick when troops departed
	ArrivalTick   int64 // Tick when troops arrive
	Resources     support.Resources
}

func (e TroopsSentEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TroopsSentEvent) AggregateID() string {
	return e.FromVillageID.String()
}

package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TrapTriggeredEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TrapTriggeredEvent{})
}

type TrapTriggeredEvent struct {
	WorldID         uuid.UUID
	VillageID       uuid.UUID
	AttackerVillage uuid.UUID
	TrappedTroops   support.Troops
	TrapsUsed       int
	TrapsRemaining  int
}

func (e TrapTriggeredEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TrapTriggeredEvent) AggregateID() string {
	return e.VillageID.String()
}

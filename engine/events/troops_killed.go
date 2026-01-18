package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TroopsKilledEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TroopsKilledEvent{})
}

type TroopsKilledEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	OwnerID   uuid.UUID
	BattleID  uuid.UUID
	Troops    support.Troops
}

func (e TroopsKilledEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TroopsKilledEvent) AggregateID() string {
	return e.VillageID.String()
}

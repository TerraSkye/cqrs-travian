package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TroopUpgradeCompletedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TroopUpgradeCompletedEvent{})
}

type TroopUpgradeCompletedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	TroopType support.TroopType
	NewLevel  int
}

func (e TroopUpgradeCompletedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TroopUpgradeCompletedEvent) AggregateID() string {
	return e.VillageID.String()
}

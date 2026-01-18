package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TroopUpgradeStartedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TroopUpgradeStartedEvent{})
}

type TroopUpgradeStartedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	TroopType support.TroopType
	FromLevel int
	ToLevel   int
	Cost      support.Resources
	StartTick int64 // Tick when upgrade started
	EndTick   int64 // Tick when upgrade completes
}

func (e TroopUpgradeStartedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TroopUpgradeStartedEvent) AggregateID() string {
	return e.VillageID.String()
}

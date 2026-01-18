package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*WallDamagedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&WallDamagedEvent{})
}

type WallDamagedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	BattleID  uuid.UUID
	OldLevel  int
	NewLevel  int
	RamsUsed  int
}

func (e WallDamagedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e WallDamagedEvent) AggregateID() string {
	return e.VillageID.String()
}

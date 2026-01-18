package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*VillageFoundedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&VillageFoundedEvent{})
}

type VillageFoundedEvent struct {
	WorldID       uuid.UUID
	PlayerID      uuid.UUID
	VillageID     uuid.UUID
	SourceVillage uuid.UUID
	Name          string
	X             int
	Y             int
	SettlersUsed  int
}

func (e VillageFoundedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e VillageFoundedEvent) AggregateID() string {
	return e.VillageID.String()
}

package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TroopTrainingQueuedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TroopTrainingQueuedEvent{})
}

type TroopTrainingQueuedEvent struct {
	WorldID       uuid.UUID
	VillageID     uuid.UUID
	QueueID       uuid.UUID
	TroopType     support.TroopType
	Amount        int
	Cost          support.Resources
	DurationTicks int64 // Duration in game ticks per unit
	BuildingSlot  int
}

func (e TroopTrainingQueuedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TroopTrainingQueuedEvent) AggregateID() string {
	return e.VillageID.String()
}

package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TroopTrainingStartedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TroopTrainingStartedEvent{})
}

type TroopTrainingStartedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	QueueID   uuid.UUID
	TroopType support.TroopType
	Amount    int
	StartTick int64 // Tick when training started
	EndTick   int64 // Tick when training completes
}

func (e TroopTrainingStartedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TroopTrainingStartedEvent) AggregateID() string {
	return e.VillageID.String()
}

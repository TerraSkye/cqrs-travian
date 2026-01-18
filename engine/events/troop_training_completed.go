package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TroopTrainingCompletedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TroopTrainingCompletedEvent{})
}

type TroopTrainingCompletedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	QueueID   uuid.UUID
	TroopType support.TroopType
	Amount    int
}

func (e TroopTrainingCompletedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TroopTrainingCompletedEvent) AggregateID() string {
	return e.VillageID.String()
}

package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*PlusAccountActivatedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&PlusAccountActivatedEvent{})
}

type PlusAccountActivatedEvent struct {
	WorldID       uuid.UUID
	PlayerID      uuid.UUID
	ActivatedTick int64 // Tick when Plus was activated
	ExpiresTick   int64 // Tick when Plus expires
	GoldSpent     int
}

func (e PlusAccountActivatedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e PlusAccountActivatedEvent) AggregateID() string {
	return e.PlayerID.String()
}

package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TickEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TickEvent{})
}

type TickEvent struct {
	WorldID   uuid.UUID
	TickCount int64
}

func (e TickEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TickEvent) AggregateID() string {
	return e.WorldID.String()
}

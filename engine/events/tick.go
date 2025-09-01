package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TickEvent)(nil)

type TickEvent struct {
	WorldID   uuid.UUID
	TickCount int64
}

func (e *TickEvent) AggregateID() string {
	return e.WorldID.String()
}

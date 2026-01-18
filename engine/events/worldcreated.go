package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*WorldCreatedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&WorldCreatedEvent{})
}

type WorldCreatedEvent struct {
	WorldID         uuid.UUID
	Name            string
	TicksPerMinute  int // Default 60 = 1 tick per second real-time
	SpeedMultiplier int // Game speed multiplier (1x, 2x, 3x, etc.)
	MapSize         int // Map radius (-MapSize to +MapSize coordinates)
}

func (e WorldCreatedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e WorldCreatedEvent) AggregateID() string {
	return e.WorldID.String()
}

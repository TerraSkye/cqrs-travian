package events

import (
	"time"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*WorldEndedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&WorldEndedEvent{})
}

type WorldEndedEvent struct {
	WorldID         uuid.UUID
	WinningAlliance uuid.UUID
	WinningVillage  uuid.UUID
	EndedAt         time.Time
}

func (e WorldEndedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e WorldEndedEvent) AggregateID() string {
	return e.WorldID.String()
}

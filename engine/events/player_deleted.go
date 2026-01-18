package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*PlayerDeletedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&PlayerDeletedEvent{})
}

type PlayerDeletedEvent struct {
	WorldID  uuid.UUID
	PlayerID uuid.UUID
	Reason   string
}

func (e PlayerDeletedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e PlayerDeletedEvent) AggregateID() string {
	return e.PlayerID.String()
}

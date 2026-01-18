package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*MessageDeletedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&MessageDeletedEvent{})
}

type MessageDeletedEvent struct {
	WorldID   uuid.UUID
	MessageID uuid.UUID
	PlayerID  uuid.UUID
}

func (e MessageDeletedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e MessageDeletedEvent) AggregateID() string {
	return e.MessageID.String()
}

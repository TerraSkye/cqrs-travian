package events

import (
	"time"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*MessageReadEvent)(nil)

func init() {
	cqrs.RegisterEvent(&MessageReadEvent{})
}

type MessageReadEvent struct {
	WorldID   uuid.UUID
	MessageID uuid.UUID
	ReaderID  uuid.UUID
	ReadAt    time.Time
}

func (e MessageReadEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e MessageReadEvent) AggregateID() string {
	return e.MessageID.String()
}

package events

import (
	"time"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*MessageSentEvent)(nil)

func init() {
	cqrs.RegisterEvent(&MessageSentEvent{})
}

type MessageSentEvent struct {
	WorldID     uuid.UUID
	MessageID   uuid.UUID
	SenderID    uuid.UUID
	RecipientID uuid.UUID
	Subject     string
	Body        string
	SentAt      time.Time
}

func (e MessageSentEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e MessageSentEvent) AggregateID() string {
	return e.MessageID.String()
}

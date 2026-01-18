package events

import (
	"time"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*IGMSentEvent)(nil)

func init() {
	cqrs.RegisterEvent(&IGMSentEvent{})
}

type IGMSentEvent struct {
	WorldID    uuid.UUID
	MessageID  uuid.UUID
	AllianceID uuid.UUID
	SenderID   uuid.UUID
	Subject    string
	Body       string
	SentAt     time.Time
}

func (e IGMSentEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e IGMSentEvent) AggregateID() string {
	return e.MessageID.String()
}

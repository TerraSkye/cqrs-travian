package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AllianceCreatedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AllianceCreatedEvent{})
}

type AllianceCreatedEvent struct {
	WorldID     uuid.UUID
	AllianceID  uuid.UUID
	FounderID   uuid.UUID
	Name        string
	Tag         string
	Description string
}

func (e AllianceCreatedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AllianceCreatedEvent) AggregateID() string {
	return e.AllianceID.String()
}

package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AllianceDescriptionUpdatedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AllianceDescriptionUpdatedEvent{})
}

type AllianceDescriptionUpdatedEvent struct {
	WorldID        uuid.UUID
	AllianceID     uuid.UUID
	OldDescription string
	NewDescription string
}

func (e AllianceDescriptionUpdatedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AllianceDescriptionUpdatedEvent) AggregateID() string {
	return e.AllianceID.String()
}

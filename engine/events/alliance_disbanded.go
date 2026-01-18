package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AllianceDisbandedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AllianceDisbandedEvent{})
}

type AllianceDisbandedEvent struct {
	WorldID    uuid.UUID
	AllianceID uuid.UUID
}

func (e AllianceDisbandedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AllianceDisbandedEvent) AggregateID() string {
	return e.AllianceID.String()
}

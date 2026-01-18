package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*CapitalChangedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&CapitalChangedEvent{})
}

type CapitalChangedEvent struct {
	WorldID      uuid.UUID
	PlayerID     uuid.UUID
	OldCapitalID uuid.UUID
	NewCapitalID uuid.UUID
}

func (e CapitalChangedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e CapitalChangedEvent) AggregateID() string {
	return e.PlayerID.String()
}

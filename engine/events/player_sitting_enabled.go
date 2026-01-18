package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*PlayerSittingEnabledEvent)(nil)

func init() {
	cqrs.RegisterEvent(&PlayerSittingEnabledEvent{})
}

type PlayerSittingEnabledEvent struct {
	WorldID  uuid.UUID
	PlayerID uuid.UUID
	SitterID uuid.UUID
}

func (e PlayerSittingEnabledEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e PlayerSittingEnabledEvent) AggregateID() string {
	return e.PlayerID.String()
}

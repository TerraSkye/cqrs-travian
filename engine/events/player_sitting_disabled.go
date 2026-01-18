package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*PlayerSittingDisabledEvent)(nil)

func init() {
	cqrs.RegisterEvent(&PlayerSittingDisabledEvent{})
}

type PlayerSittingDisabledEvent struct {
	WorldID  uuid.UUID
	PlayerID uuid.UUID
	SitterID uuid.UUID
}

func (e PlayerSittingDisabledEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e PlayerSittingDisabledEvent) AggregateID() string {
	return e.PlayerID.String()
}

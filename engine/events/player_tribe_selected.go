package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*PlayerTribeSelectedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&PlayerTribeSelectedEvent{})
}

type PlayerTribeSelectedEvent struct {
	WorldID  uuid.UUID
	PlayerID uuid.UUID
	Tribe    support.Tribe
}

func (e PlayerTribeSelectedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e PlayerTribeSelectedEvent) AggregateID() string {
	return e.PlayerID.String()
}

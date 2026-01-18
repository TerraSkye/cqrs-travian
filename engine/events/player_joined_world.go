package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*PlayerJoinedWorldEvent)(nil)

func init() {
	cqrs.RegisterEvent(&PlayerJoinedWorldEvent{})
}

type PlayerJoinedWorldEvent struct {
	WorldID   uuid.UUID
	PlayerID  uuid.UUID
	AccountID uuid.UUID
	Username  string
	Tribe     support.Tribe
	StartX    int
	StartY    int
}

func (e PlayerJoinedWorldEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e PlayerJoinedWorldEvent) AggregateID() string {
	return e.PlayerID.String()
}

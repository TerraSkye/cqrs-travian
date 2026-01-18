package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AllianceMemberJoinedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AllianceMemberJoinedEvent{})
}

type AllianceMemberJoinedEvent struct {
	WorldID    uuid.UUID
	AllianceID uuid.UUID
	PlayerID   uuid.UUID
	Role       support.AllianceRole
}

func (e AllianceMemberJoinedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AllianceMemberJoinedEvent) AggregateID() string {
	return e.AllianceID.String()
}

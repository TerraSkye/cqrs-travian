package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AllianceMemberLeftEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AllianceMemberLeftEvent{})
}

type AllianceMemberLeftEvent struct {
	WorldID    uuid.UUID
	AllianceID uuid.UUID
	PlayerID   uuid.UUID
}

func (e AllianceMemberLeftEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AllianceMemberLeftEvent) AggregateID() string {
	return e.AllianceID.String()
}

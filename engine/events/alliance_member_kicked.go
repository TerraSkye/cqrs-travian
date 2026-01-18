package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AllianceMemberKickedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AllianceMemberKickedEvent{})
}

type AllianceMemberKickedEvent struct {
	WorldID    uuid.UUID
	AllianceID uuid.UUID
	PlayerID   uuid.UUID
	KickedByID uuid.UUID
	Reason     string
}

func (e AllianceMemberKickedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AllianceMemberKickedEvent) AggregateID() string {
	return e.AllianceID.String()
}

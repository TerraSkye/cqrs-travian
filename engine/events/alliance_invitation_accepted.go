package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AllianceInvitationAcceptedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AllianceInvitationAcceptedEvent{})
}

type AllianceInvitationAcceptedEvent struct {
	WorldID      uuid.UUID
	AllianceID   uuid.UUID
	InvitationID uuid.UUID
	PlayerID     uuid.UUID
}

func (e AllianceInvitationAcceptedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AllianceInvitationAcceptedEvent) AggregateID() string {
	return e.AllianceID.String()
}

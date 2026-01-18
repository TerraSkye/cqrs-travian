package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AllianceInvitationDeclinedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AllianceInvitationDeclinedEvent{})
}

type AllianceInvitationDeclinedEvent struct {
	WorldID      uuid.UUID
	AllianceID   uuid.UUID
	InvitationID uuid.UUID
	PlayerID     uuid.UUID
}

func (e AllianceInvitationDeclinedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AllianceInvitationDeclinedEvent) AggregateID() string {
	return e.AllianceID.String()
}

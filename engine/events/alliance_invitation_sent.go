package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AllianceInvitationSentEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AllianceInvitationSentEvent{})
}

type AllianceInvitationSentEvent struct {
	WorldID      uuid.UUID
	AllianceID   uuid.UUID
	InvitationID uuid.UUID
	InviterID    uuid.UUID
	InviteeID    uuid.UUID
	SentTick     int64 // Tick when invitation was sent
	ExpiresTick  int64 // Tick when invitation expires
}

func (e AllianceInvitationSentEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AllianceInvitationSentEvent) AggregateID() string {
	return e.AllianceID.String()
}

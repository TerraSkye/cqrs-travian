package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AllianceRoleChangedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AllianceRoleChangedEvent{})
}

type AllianceRoleChangedEvent struct {
	WorldID    uuid.UUID
	AllianceID uuid.UUID
	PlayerID   uuid.UUID
	OldRole    support.AllianceRole
	NewRole    support.AllianceRole
	ChangedBy  uuid.UUID
}

func (e AllianceRoleChangedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AllianceRoleChangedEvent) AggregateID() string {
	return e.AllianceID.String()
}

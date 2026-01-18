package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*VillageAbandonedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&VillageAbandonedEvent{})
}

type VillageAbandonedEvent struct {
	WorldID   uuid.UUID
	PlayerID  uuid.UUID
	VillageID uuid.UUID
}

func (e VillageAbandonedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e VillageAbandonedEvent) AggregateID() string {
	return e.VillageID.String()
}

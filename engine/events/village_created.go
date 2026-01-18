package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*VillageCreatedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&VillageCreatedEvent{})
}

type VillageCreatedEvent struct {
	WorldID   uuid.UUID
	PlayerID  uuid.UUID
	VillageID uuid.UUID
	Name      string
	X         int
	Y         int
	Tribe     support.Tribe
	IsCapital bool
}

func (e VillageCreatedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e VillageCreatedEvent) AggregateID() string {
	return e.VillageID.String()
}

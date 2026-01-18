package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*HeroCreatedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&HeroCreatedEvent{})
}

type HeroCreatedEvent struct {
	WorldID   uuid.UUID
	PlayerID  uuid.UUID
	HeroID    uuid.UUID
	VillageID uuid.UUID
	Tribe     support.Tribe
}

func (e HeroCreatedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e HeroCreatedEvent) AggregateID() string {
	return e.HeroID.String()
}

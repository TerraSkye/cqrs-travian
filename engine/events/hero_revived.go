package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*HeroRevivedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&HeroRevivedEvent{})
}

type HeroRevivedEvent struct {
	WorldID   uuid.UUID
	HeroID    uuid.UUID
	PlayerID  uuid.UUID
	VillageID uuid.UUID
	Health    int
}

func (e HeroRevivedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e HeroRevivedEvent) AggregateID() string {
	return e.HeroID.String()
}

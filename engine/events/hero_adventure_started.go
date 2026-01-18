package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*HeroAdventureStartedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&HeroAdventureStartedEvent{})
}

type HeroAdventureStartedEvent struct {
	WorldID       uuid.UUID
	HeroID        uuid.UUID
	AdventureID   uuid.UUID
	X             int
	Y             int
	Difficulty    int
	DepartureTick int64 // Tick when hero departed
	ArrivalTick   int64 // Tick when hero arrives at adventure
}

func (e HeroAdventureStartedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e HeroAdventureStartedEvent) AggregateID() string {
	return e.HeroID.String()
}

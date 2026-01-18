package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AdventureSpawnedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AdventureSpawnedEvent{})
}

type AdventureSpawnedEvent struct {
	WorldID     uuid.UUID
	AdventureID uuid.UUID
	PlayerID    uuid.UUID
	X           int
	Y           int
	Difficulty  int
	SpawnedTick int64 // Tick when adventure spawned
	ExpiresTick int64 // Tick when adventure expires
}

func (e AdventureSpawnedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AdventureSpawnedEvent) AggregateID() string {
	return e.AdventureID.String()
}

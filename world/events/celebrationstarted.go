package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*CelebrationStartedEvent)(nil)

type CelebrationStartedEvent struct {
	CelebrationID uuid.UUID
	WorldID       uuid.UUID
	VillageID     uuid.UUID
	Large         bool
	// duration of the celebration
	Duration int64
}

func (e *CelebrationStartedEvent) AggregateID() uuid.UUID {
	return e.VillageID
}

package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*CelebrationCompletedEvent)(nil)

type CelebrationCompletedEvent struct {
	CelebrationID uuid.UUID
	WorldID       uuid.UUID
	VillageID     uuid.UUID
	// culture points produced
	CulturePoints int64
}

func (e *CelebrationCompletedEvent) AggregateID() uuid.UUID {
	return e.VillageID
}

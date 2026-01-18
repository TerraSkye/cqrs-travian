package events

import (
	"time"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*WorldWonderCompletedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&WorldWonderCompletedEvent{})
}

type WorldWonderCompletedEvent struct {
	WorldID     uuid.UUID
	VillageID   uuid.UUID
	AllianceID  uuid.UUID
	CompletedAt time.Time
}

func (e WorldWonderCompletedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e WorldWonderCompletedEvent) AggregateID() string {
	return e.VillageID.String()
}

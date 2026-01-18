package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*WorldWonderDestroyedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&WorldWonderDestroyedEvent{})
}

type WorldWonderDestroyedEvent struct {
	WorldID        uuid.UUID
	VillageID      uuid.UUID
	AllianceID     uuid.UUID
	AttackerPlayer uuid.UUID
	OldLevel       int
	NewLevel       int
}

func (e WorldWonderDestroyedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e WorldWonderDestroyedEvent) AggregateID() string {
	return e.VillageID.String()
}

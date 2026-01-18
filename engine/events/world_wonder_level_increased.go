package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*WorldWonderLevelIncreasedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&WorldWonderLevelIncreasedEvent{})
}

type WorldWonderLevelIncreasedEvent struct {
	WorldID    uuid.UUID
	VillageID  uuid.UUID
	AllianceID uuid.UUID
	OldLevel   int
	NewLevel   int
}

func (e WorldWonderLevelIncreasedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e WorldWonderLevelIncreasedEvent) AggregateID() string {
	return e.VillageID.String()
}

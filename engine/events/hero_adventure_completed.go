package events

import (
	support2 "cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*HeroAdventureCompletedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&HeroAdventureCompletedEvent{})
}

type HeroAdventureCompletedEvent struct {
	WorldID        uuid.UUID
	HeroID         uuid.UUID
	AdventureID    uuid.UUID
	Success        bool
	ExperienceGain int
	ItemFound      support2.ItemType
	ResourcesFound support2.Resources
	HealthLost     int
}

func (e HeroAdventureCompletedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e HeroAdventureCompletedEvent) AggregateID() string {
	return e.HeroID.String()
}

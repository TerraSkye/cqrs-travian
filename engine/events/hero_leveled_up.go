package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*HeroLeveledUpEvent)(nil)

func init() {
	cqrs.RegisterEvent(&HeroLeveledUpEvent{})
}

type HeroLeveledUpEvent struct {
	WorldID     uuid.UUID
	HeroID      uuid.UUID
	PlayerID    uuid.UUID
	OldLevel    int
	NewLevel    int
	SkillPoints int
}

func (e HeroLeveledUpEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e HeroLeveledUpEvent) AggregateID() string {
	return e.HeroID.String()
}

package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*HeroHealthChangedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&HeroHealthChangedEvent{})
}

type HeroHealthChangedEvent struct {
	WorldID   uuid.UUID
	HeroID    uuid.UUID
	OldHealth int
	NewHealth int
	Reason    string
}

func (e HeroHealthChangedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e HeroHealthChangedEvent) AggregateID() string {
	return e.HeroID.String()
}

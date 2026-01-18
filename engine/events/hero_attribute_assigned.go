package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*HeroAttributeAssignedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&HeroAttributeAssignedEvent{})
}

type HeroAttributeAssignedEvent struct {
	WorldID   uuid.UUID
	HeroID    uuid.UUID
	Attribute support.HeroAttribute
	Points    int
}

func (e HeroAttributeAssignedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e HeroAttributeAssignedEvent) AggregateID() string {
	return e.HeroID.String()
}

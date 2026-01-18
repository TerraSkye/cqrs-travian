package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*HeroEquipmentChangedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&HeroEquipmentChangedEvent{})
}

type HeroEquipmentChangedEvent struct {
	WorldID uuid.UUID
	HeroID  uuid.UUID
	Slot    support.EquipmentSlot
	OldItem support.ItemType
	NewItem support.ItemType
}

func (e HeroEquipmentChangedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e HeroEquipmentChangedEvent) AggregateID() string {
	return e.HeroID.String()
}

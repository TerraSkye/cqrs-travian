package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*NatarVillageSpawnedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&NatarVillageSpawnedEvent{})
}

type NatarVillageSpawnedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	Name      string
	X         int
	Y         int
	Size      string
}

func (e NatarVillageSpawnedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e NatarVillageSpawnedEvent) AggregateID() string {
	return e.VillageID.String()
}

package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TrapReleasedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TrapReleasedEvent{})
}

type TrapReleasedEvent struct {
	WorldID        uuid.UUID
	VillageID      uuid.UUID
	OwnerVillage   uuid.UUID
	ReleasedTroops support.Troops
	ReleaseType    string
}

func (e TrapReleasedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TrapReleasedEvent) AggregateID() string {
	return e.VillageID.String()
}

package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TroopsDisbandedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TroopsDisbandedEvent{})
}

type TroopsDisbandedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	Troops    support.Troops
}

func (e TroopsDisbandedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TroopsDisbandedEvent) AggregateID() string {
	return e.VillageID.String()
}

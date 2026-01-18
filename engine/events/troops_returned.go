package events

import (
	support2 "cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TroopsReturnedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TroopsReturnedEvent{})
}

type TroopsReturnedEvent struct {
	WorldID    uuid.UUID
	MovementID uuid.UUID
	VillageID  uuid.UUID
	Troops     support2.Troops
	Resources  support2.Resources
}

func (e TroopsReturnedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TroopsReturnedEvent) AggregateID() string {
	return e.VillageID.String()
}

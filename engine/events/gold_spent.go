package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*GoldSpentEvent)(nil)

func init() {
	cqrs.RegisterEvent(&GoldSpentEvent{})
}

type GoldSpentEvent struct {
	WorldID   uuid.UUID
	PlayerID  uuid.UUID
	VillageID uuid.UUID
	Amount    int
	Purpose   string
}

func (e GoldSpentEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e GoldSpentEvent) AggregateID() string {
	return e.PlayerID.String()
}

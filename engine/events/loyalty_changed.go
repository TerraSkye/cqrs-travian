package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*LoyaltyChangedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&LoyaltyChangedEvent{})
}

type LoyaltyChangedEvent struct {
	WorldID    uuid.UUID
	VillageID  uuid.UUID
	OldLoyalty int
	NewLoyalty int
	AttackerID uuid.UUID
}

func (e LoyaltyChangedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e LoyaltyChangedEvent) AggregateID() string {
	return e.VillageID.String()
}

package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*VillageConqueredEvent)(nil)

func init() {
	cqrs.RegisterEvent(&VillageConqueredEvent{})
}

type VillageConqueredEvent struct {
	WorldID          uuid.UUID
	VillageID        uuid.UUID
	PreviousOwnerID  uuid.UUID
	NewOwnerID       uuid.UUID
	LoyaltyReduction int
}

func (e VillageConqueredEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e VillageConqueredEvent) AggregateID() string {
	return e.VillageID.String()
}

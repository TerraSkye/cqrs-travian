package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ResourcesTransferredEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ResourcesTransferredEvent{})
}

type ResourcesTransferredEvent struct {
	WorldID       uuid.UUID
	FromVillageID uuid.UUID
	ToVillageID   uuid.UUID
	Amount        support.Resources
	MerchantCount int
}

func (e ResourcesTransferredEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ResourcesTransferredEvent) AggregateID() string {
	return e.FromVillageID.String()
}

package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*MerchantsReturnedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&MerchantsReturnedEvent{})
}

type MerchantsReturnedEvent struct {
	WorldID       uuid.UUID
	TradeID       uuid.UUID
	VillageID     uuid.UUID
	MerchantCount int
}

func (e MerchantsReturnedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e MerchantsReturnedEvent) AggregateID() string {
	return e.TradeID.String()
}

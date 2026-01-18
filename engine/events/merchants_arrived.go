package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*MerchantsArrivedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&MerchantsArrivedEvent{})
}

type MerchantsArrivedEvent struct {
	WorldID       uuid.UUID
	TradeID       uuid.UUID
	FromVillageID uuid.UUID
	ToVillageID   uuid.UUID
	Resources     support.Resources
	MerchantCount int
}

func (e MerchantsArrivedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e MerchantsArrivedEvent) AggregateID() string {
	return e.TradeID.String()
}

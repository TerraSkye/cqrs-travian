package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TradeOfferCreatedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TradeOfferCreatedEvent{})
}

type TradeOfferCreatedEvent struct {
	WorldID       uuid.UUID
	OfferID       uuid.UUID
	VillageID     uuid.UUID
	Offering      support.Resources
	Seeking       support.Resources
	MerchantCount int
	CreatedTick   int64 // Tick when offer was created
	ExpiresTick   int64 // Tick when offer expires
	MaxRange      int
}

func (e TradeOfferCreatedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TradeOfferCreatedEvent) AggregateID() string {
	return e.OfferID.String()
}

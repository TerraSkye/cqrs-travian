package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TradeOfferCancelledEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TradeOfferCancelledEvent{})
}

type TradeOfferCancelledEvent struct {
	WorldID   uuid.UUID
	OfferID   uuid.UUID
	VillageID uuid.UUID
}

func (e TradeOfferCancelledEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TradeOfferCancelledEvent) AggregateID() string {
	return e.OfferID.String()
}

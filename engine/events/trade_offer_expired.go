package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TradeOfferExpiredEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TradeOfferExpiredEvent{})
}

type TradeOfferExpiredEvent struct {
	WorldID   uuid.UUID
	OfferID   uuid.UUID
	VillageID uuid.UUID
}

func (e TradeOfferExpiredEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TradeOfferExpiredEvent) AggregateID() string {
	return e.OfferID.String()
}

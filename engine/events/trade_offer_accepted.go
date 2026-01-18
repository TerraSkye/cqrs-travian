package events

import (
	"time"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*TradeOfferAcceptedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&TradeOfferAcceptedEvent{})
}

type TradeOfferAcceptedEvent struct {
	WorldID          uuid.UUID
	OfferID          uuid.UUID
	AcceptingVillage uuid.UUID
	OfferingVillage  uuid.UUID
	DeliveryTime     time.Time
}

func (e TradeOfferAcceptedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e TradeOfferAcceptedEvent) AggregateID() string {
	return e.OfferID.String()
}

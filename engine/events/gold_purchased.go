package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*GoldPurchasedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&GoldPurchasedEvent{})
}

type GoldPurchasedEvent struct {
	WorldID       uuid.UUID
	PlayerID      uuid.UUID
	Amount        int
	TransactionID string
}

func (e GoldPurchasedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e GoldPurchasedEvent) AggregateID() string {
	return e.PlayerID.String()
}

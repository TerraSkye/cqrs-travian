package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*MerchantsSentEvent)(nil)

func init() {
	cqrs.RegisterEvent(&MerchantsSentEvent{})
}

type MerchantsSentEvent struct {
	WorldID       uuid.UUID
	TradeID       uuid.UUID
	FromVillageID uuid.UUID
	ToVillageID   uuid.UUID
	Resources     support.Resources
	MerchantCount int
	DepartureTick int64 // Tick when merchants departed
	ArrivalTick   int64 // Tick when merchants arrive
}

func (e MerchantsSentEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e MerchantsSentEvent) AggregateID() string {
	return e.TradeID.String()
}

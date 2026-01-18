package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*GoldCompletionEvent)(nil)

func init() {
	cqrs.RegisterEvent(&GoldCompletionEvent{})
}

type GoldCompletionEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
}

func (w GoldCompletionEvent) EventType() string {
	return cqrs.TypeName(w)
}

func (w GoldCompletionEvent) AggregateID() string {
	return w.WorldID.String()
}

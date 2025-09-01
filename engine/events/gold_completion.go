package events

import (
	"github.com/google/uuid"
)

type GoldCompletionEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
}

func (w GoldCompletionEvent) AggregateID() string {
	return w.WorldID.String()
}

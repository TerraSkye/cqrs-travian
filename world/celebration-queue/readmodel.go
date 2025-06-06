package celebration_queue

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.ReadModel = (*ReadModel)(nil)

type Query struct {
}

type ReadModel struct {
	AggregateId uuid.UUID
	VillageID   uuid.UUID
	WorldID     uuid.UUID
	TTL         int64
	Large       bool
}

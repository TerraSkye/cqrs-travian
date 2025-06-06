package villageoverview

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
	Coordinate  struct {
		X int64
		T int64
	}
	Movement struct {
		Type int8
		TTL  int64
	}
	Production struct {
		Wood int64
		Crop int64
		Clay int64
		Iron int64
	}
	Loyalty   int64
	IsCapital int64

	Queue []struct {
		Upgrade string
		TTL     int64
	}
}

package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*BuildingDemolishedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&BuildingDemolishedEvent{})
}

type BuildingDemolishedEvent struct {
	WorldID      uuid.UUID
	VillageID    uuid.UUID
	BuildingType support.BuildingType
	Slot         int
	FromLevel    int
	ToLevel      int
}

func (e BuildingDemolishedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e BuildingDemolishedEvent) AggregateID() string {
	return e.VillageID.String()
}

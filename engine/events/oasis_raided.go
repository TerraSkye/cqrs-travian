package events

import (
	support2 "cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*OasisRaidedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&OasisRaidedEvent{})
}

type OasisRaidedEvent struct {
	WorldID         uuid.UUID
	OasisID         uuid.UUID
	AttackerVillage uuid.UUID
	AnimalsKilled   support2.Troops
	ResourcesLooted support2.Resources
}

func (e OasisRaidedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e OasisRaidedEvent) AggregateID() string {
	return e.OasisID.String()
}

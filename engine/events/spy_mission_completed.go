package events

import (
	support2 "cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*SpyMissionCompletedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&SpyMissionCompletedEvent{})
}

type SpyMissionCompletedEvent struct {
	WorldID         uuid.UUID
	MissionID       uuid.UUID
	AttackerVillage uuid.UUID
	DefenderVillage uuid.UUID
	Success         bool
	SpiesSent       int
	SpiesLost       int
	DefenderSpies   int
	ResourcesFound  support2.Resources
	TroopsFound     support2.Troops
	DefensesFound   bool
}

func (e SpyMissionCompletedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e SpyMissionCompletedEvent) AggregateID() string {
	return e.MissionID.String()
}

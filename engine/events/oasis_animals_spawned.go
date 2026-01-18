package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*OasisAnimalsSpawnedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&OasisAnimalsSpawnedEvent{})
}

type OasisAnimalsSpawnedEvent struct {
	WorldID uuid.UUID
	OasisID uuid.UUID
	Animals support.Troops
}

func (e OasisAnimalsSpawnedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e OasisAnimalsSpawnedEvent) AggregateID() string {
	return e.OasisID.String()
}

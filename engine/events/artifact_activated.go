package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ArtifactActivatedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ArtifactActivatedEvent{})
}

type ArtifactActivatedEvent struct {
	WorldID       uuid.UUID
	ArtifactID    uuid.UUID
	VillageID     uuid.UUID
	PlayerID      uuid.UUID
	ActivatedTick int64 // Tick when artifact was activated
}

func (e ArtifactActivatedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ArtifactActivatedEvent) AggregateID() string {
	return e.ArtifactID.String()
}

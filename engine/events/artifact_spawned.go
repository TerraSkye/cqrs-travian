package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ArtifactSpawnedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ArtifactSpawnedEvent{})
}

type ArtifactSpawnedEvent struct {
	WorldID      uuid.UUID
	ArtifactID   uuid.UUID
	VillageID    uuid.UUID
	ArtifactType string
	Size         string
	X            int
	Y            int
}

func (e ArtifactSpawnedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ArtifactSpawnedEvent) AggregateID() string {
	return e.ArtifactID.String()
}

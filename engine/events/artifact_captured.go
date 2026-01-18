package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ArtifactCapturedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ArtifactCapturedEvent{})
}

type ArtifactCapturedEvent struct {
	WorldID         uuid.UUID
	ArtifactID      uuid.UUID
	PreviousVillage uuid.UUID
	PreviousPlayer  uuid.UUID
	NewVillage      uuid.UUID
	NewPlayer       uuid.UUID
}

func (e ArtifactCapturedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ArtifactCapturedEvent) AggregateID() string {
	return e.ArtifactID.String()
}

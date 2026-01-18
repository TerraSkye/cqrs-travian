package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*OasisLostEvent)(nil)

func init() {
	cqrs.RegisterEvent(&OasisLostEvent{})
}

type OasisLostEvent struct {
	WorldID         uuid.UUID
	OasisID         uuid.UUID
	PreviousVillage uuid.UUID
	PreviousPlayer  uuid.UUID
	NewVillage      uuid.UUID
	NewPlayer       uuid.UUID
}

func (e OasisLostEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e OasisLostEvent) AggregateID() string {
	return e.OasisID.String()
}

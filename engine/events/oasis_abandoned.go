package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*OasisAbandonedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&OasisAbandonedEvent{})
}

type OasisAbandonedEvent struct {
	WorldID   uuid.UUID
	OasisID   uuid.UUID
	VillageID uuid.UUID
	PlayerID  uuid.UUID
}

func (e OasisAbandonedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e OasisAbandonedEvent) AggregateID() string {
	return e.OasisID.String()
}

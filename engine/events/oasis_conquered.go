package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*OasisConqueredEvent)(nil)

func init() {
	cqrs.RegisterEvent(&OasisConqueredEvent{})
}

type OasisConqueredEvent struct {
	WorldID   uuid.UUID
	OasisID   uuid.UUID
	VillageID uuid.UUID
	PlayerID  uuid.UUID
	OasisType support.OasisType
	X         int
	Y         int
}

func (e OasisConqueredEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e OasisConqueredEvent) AggregateID() string {
	return e.OasisID.String()
}

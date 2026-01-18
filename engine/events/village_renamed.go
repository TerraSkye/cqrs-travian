package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*VillageRenamedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&VillageRenamedEvent{})
}

type VillageRenamedEvent struct {
	WorldID   uuid.UUID
	VillageID uuid.UUID
	OldName   string
	NewName   string
}

func (e VillageRenamedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e VillageRenamedEvent) AggregateID() string {
	return e.VillageID.String()
}

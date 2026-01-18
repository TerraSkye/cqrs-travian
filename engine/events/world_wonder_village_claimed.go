package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*WorldWonderVillageClaimedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&WorldWonderVillageClaimedEvent{})
}

type WorldWonderVillageClaimedEvent struct {
	WorldID    uuid.UUID
	VillageID  uuid.UUID
	PlayerID   uuid.UUID
	AllianceID uuid.UUID
	X          int
	Y          int
}

func (e WorldWonderVillageClaimedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e WorldWonderVillageClaimedEvent) AggregateID() string {
	return e.VillageID.String()
}

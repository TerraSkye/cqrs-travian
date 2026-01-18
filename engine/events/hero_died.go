package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*HeroDiedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&HeroDiedEvent{})
}

type HeroDiedEvent struct {
	WorldID  uuid.UUID
	HeroID   uuid.UUID
	PlayerID uuid.UUID
	BattleID uuid.UUID
	Cause    string
}

func (e HeroDiedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e HeroDiedEvent) AggregateID() string {
	return e.HeroID.String()
}

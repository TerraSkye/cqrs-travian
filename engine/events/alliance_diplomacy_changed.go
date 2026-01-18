package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AllianceDiplomacyChangedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AllianceDiplomacyChangedEvent{})
}

type AllianceDiplomacyChangedEvent struct {
	WorldID          uuid.UUID
	AllianceID       uuid.UUID
	TargetAllianceID uuid.UUID
	OldRelation      support.DiplomacyType
	NewRelation      support.DiplomacyType
}

func (e AllianceDiplomacyChangedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AllianceDiplomacyChangedEvent) AggregateID() string {
	return e.AllianceID.String()
}

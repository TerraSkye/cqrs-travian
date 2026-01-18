package events

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ReportDeletedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ReportDeletedEvent{})
}

type ReportDeletedEvent struct {
	WorldID  uuid.UUID
	ReportID uuid.UUID
	PlayerID uuid.UUID
}

func (e ReportDeletedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ReportDeletedEvent) AggregateID() string {
	return e.ReportID.String()
}

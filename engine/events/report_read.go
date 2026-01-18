package events

import (
	"time"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ReportReadEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ReportReadEvent{})
}

type ReportReadEvent struct {
	WorldID  uuid.UUID
	ReportID uuid.UUID
	PlayerID uuid.UUID
	ReadAt   time.Time
}

func (e ReportReadEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ReportReadEvent) AggregateID() string {
	return e.ReportID.String()
}

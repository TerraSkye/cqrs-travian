package events

import (
	"cqrs-travian/engine/support"
	"time"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*ReportCreatedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&ReportCreatedEvent{})
}

type ReportCreatedEvent struct {
	WorldID    uuid.UUID
	ReportID   uuid.UUID
	PlayerID   uuid.UUID
	ReportType support.ReportType
	Title      string
	RelatedID  uuid.UUID
	CreatedAt  time.Time
}

func (e ReportCreatedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e ReportCreatedEvent) AggregateID() string {
	return e.ReportID.String()
}

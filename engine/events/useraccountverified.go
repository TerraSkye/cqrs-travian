package events

import (
	"fmt"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*UserAccountVerifiedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&UserAccountVerifiedEvent{})
}

type UserAccountVerifiedEvent struct {
	WorldID  uuid.UUID
	Username string
}

func (e UserAccountVerifiedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e UserAccountVerifiedEvent) AggregateID() string {
	return fmt.Sprintf("%s-%s", e.WorldID.String(), e.Username)
}

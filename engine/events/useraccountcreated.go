package events

import (
	"fmt"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*UserAccountCreatedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&UserAccountCreatedEvent{})
}

type UserAccountCreatedEvent struct {
	WorldID  uuid.UUID
	Username string
	Email    string
	Password string
}

func (e UserAccountCreatedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e UserAccountCreatedEvent) AggregateID() string {
	return fmt.Sprintf("%s-%s", e.WorldID.String(), e.Username)
}

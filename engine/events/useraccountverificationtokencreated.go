package events

import (
	"fmt"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*UserAccountVerificationTokenCreatedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&UserAccountVerificationTokenCreatedEvent{})
}

type UserAccountVerificationTokenCreatedEvent struct {
	WorldID  uuid.UUID
	Username string
	Email    string
	Token    string
}

func (e UserAccountVerificationTokenCreatedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e UserAccountVerificationTokenCreatedEvent) AggregateID() string {
	return fmt.Sprintf("%s-%s", e.WorldID.String(), e.Username)
}

package events

import (
	"fmt"
	"github.com/google/uuid"
)

type UserAccountCreatedEvent struct {
	WorldID  uuid.UUID
	Username string
	Email    string
	Password string
}

func (w UserAccountCreatedEvent) AggregateID() string {
	return fmt.Sprintf("%s-%s", w.WorldID.String(), w.Username)
}

package events

import (
	"fmt"
	"github.com/google/uuid"
)

type UserAccountVerificationTokenCreated struct {
	WorldID  uuid.UUID
	Username string
	Email    string
	Token    string
}

func (w UserAccountVerificationTokenCreated) AggregateID() string {
	return fmt.Sprintf("%s-%s", w.WorldID.String(), w.Username)
}

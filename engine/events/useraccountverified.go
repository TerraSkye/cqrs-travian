package events

import (
	"fmt"
	"github.com/google/uuid"
)

type UserAccountVerified struct {
	WorldID  uuid.UUID
	Username string
}

func (w UserAccountVerified) AggregateID() string {
	return fmt.Sprintf("%s-%s", w.WorldID.String(), w.Username)
}

package commands

import (
	"github.com/google/uuid"
)

type CreateWorldCommand struct {
	WorldID uuid.UUID
	Name    string
}

func (e *CreateWorldCommand) AggregateID() string {
	return e.WorldID.String()
}

package commands

import (
	"github.com/google/uuid"
)

type Tick struct {
	WorldID uuid.UUID
}

func (e *Tick) AggregateID() string {
	return e.WorldID.String()
}

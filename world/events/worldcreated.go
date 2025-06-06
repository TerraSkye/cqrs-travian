package events

import "github.com/google/uuid"

type WorldCreatedEvent struct {
	WorldID uuid.UUID
	Name    string
}

func (w WorldCreatedEvent) AggregateID() string {
	return w.WorldID.String()
}

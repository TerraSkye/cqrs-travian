package events

import (
	"github.com/google/uuid"
)

var _ infra.Event[string] = (*TaskCompletedEvent)(nil)

type TaskCompletedEvent struct {
	WorldID uuid.UUID
	Owner   string
}

func (e *TaskCompletedEvent) AggregateID() string {
	return e.Owner
}

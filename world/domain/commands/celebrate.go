package commands

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Command = (*StartCelebrationCommand)(nil)

type StartCelebrationCommand struct {
	WorldID   uuid.UUID
	VillageID int64
	Large     bool
}

func (e *StartCelebrationCommand) AggregateID() string {
	return e.WorldID.String()
}

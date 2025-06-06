package commands

import (
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Command = (*EstablishVillageCommand)(nil)

type EstablishVillageCommand struct {
	WorldID uuid.UUID
	TileID  int64
	Large   bool
}

func (e *EstablishVillageCommand) AggregateID() uuid.UUID {
	return e.WorldID
}

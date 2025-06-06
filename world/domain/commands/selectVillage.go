package commands

import "github.com/google/uuid"

type SelectVillageCommand struct {
	WorldID   uuid.UUID
	VillageID int64
}

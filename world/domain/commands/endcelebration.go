package commands

import "github.com/google/uuid"

type EndCelebrationCommand struct {
	CelebrationID uuid.UUID
	WorldID       uuid.UUID
	VillageID     int64
}

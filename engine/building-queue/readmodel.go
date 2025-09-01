package building_queue

import "github.com/google/uuid"

type PendingConstruction struct {
	WorldID        uuid.UUID
	VillageID      uuid.UUID
	ConstructionID uuid.UUID
}

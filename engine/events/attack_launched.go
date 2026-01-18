package events

import (
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*AttackLaunchedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&AttackLaunchedEvent{})
}

type AttackLaunchedEvent struct {
	WorldID         uuid.UUID
	AttackID        uuid.UUID
	AttackerVillage uuid.UUID
	AttackerPlayer  uuid.UUID
	TargetX         int
	TargetY         int
	TargetVillage   uuid.UUID
	Troops          support.Troops
	AttackType      support.MovementType
	DepartureTick   int64 // Tick when attack was launched
	ArrivalTick     int64 // Tick when attack arrives
	CatapultTarget1 support.BuildingType
	CatapultTarget2 support.BuildingType
}

func (e AttackLaunchedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e AttackLaunchedEvent) AggregateID() string {
	return e.AttackID.String()
}

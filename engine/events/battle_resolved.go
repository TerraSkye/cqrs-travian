package events

import (
	support2 "cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Event = (*BattleResolvedEvent)(nil)

func init() {
	cqrs.RegisterEvent(&BattleResolvedEvent{})
}

type BattleResolvedEvent struct {
	WorldID              uuid.UUID
	BattleID             uuid.UUID
	AttackID             uuid.UUID
	AttackerVillage      uuid.UUID
	DefenderVillage      uuid.UUID
	AttackerTroopsBefore support2.Troops
	AttackerTroopsAfter  support2.Troops
	DefenderTroopsBefore support2.Troops
	DefenderTroopsAfter  support2.Troops
	AttackerWon          bool
	ResourcesLooted      support2.Resources
	WallDamage           int
	BuildingDamage       map[int]int
	LoyaltyChange        int
}

func (e BattleResolvedEvent) EventType() string {
	return cqrs.TypeName(e)
}

func (e BattleResolvedEvent) AggregateID() string {
	return e.BattleID.String()
}

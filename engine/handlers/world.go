package handlers

import (
	"cqrs-travian/engine"
	"cqrs-travian/engine/domain"
	cqrs "github.com/terraskye/eventsourcing"
)

func init() {
	engine.RegisterAggregate(func(id string) *domain.World {
		return &domain.World{
			AggregateBase: cqrs.NewAggregateBase(id),
		}
	})
}

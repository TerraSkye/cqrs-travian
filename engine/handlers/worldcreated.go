package handlers

import (
	"cqrs-travian/engine"
	"cqrs-travian/engine/domain"
	"cqrs-travian/engine/events"
)

func init() {
	engine.RegisterEvent(func(aggregate *domain.World) func(event *events.WorldCreatedEvent) {
		return aggregate.OnWorldCreated
	})
}

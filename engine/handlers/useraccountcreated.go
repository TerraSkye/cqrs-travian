package handlers

import (
	"cqrs-travian/engine"
	"cqrs-travian/engine/domain"
	"cqrs-travian/engine/events"
)

func init() {
	engine.RegisterEvent(func(aggregate *domain.Account) func(event *events.UserAccountCreatedEvent) {
		return aggregate.OnUserAccountCreated
	})
}

package unverified_accounts

import (
	"context"
	"cqrs-travian/engine/events"

	cqrs "github.com/terraskye/eventsourcing"
)

func NewProjector() *cqrs.EventGroupProcessor {

	p := &service{}

	return cqrs.NewEventGroupProcessor(
		"unverified-accounts",
		cqrs.NewGroupEventHandler(p.OnWorldCreatedEvent),                   // generates an overview entry for the worlds
		cqrs.NewGroupEventHandler(p.OnUserAccountVerificationTokenCreated), // generates an entry for an account within a engine
		cqrs.NewGroupEventHandler(p.OnUserAccountVerified),                 // removes the entry for an account within a engine.
	)
}

type service struct {
	worldInfo map[string]string
}

func (s *service) OnWorldCreatedEvent(ctx context.Context, ev *events.WorldCreatedEvent) error {
	s.worldInfo[ev.WorldID.String()] = ev.Name
	return nil
}

func (s *service) OnUserAccountVerificationTokenCreated(ctx context.Context, ev *events.UserAccountVerificationTokenCreated) error {
	// append to todo list
	//ev.Token
	//ev.Username
	//ev.Email
	return nil
}

func (s *service) OnUserAccountVerified(ctx context.Context, ev *events.UserAccountVerified) error {
	//remove from todo-list

	//ev.Username
	//ev.WorldID
	return nil
}

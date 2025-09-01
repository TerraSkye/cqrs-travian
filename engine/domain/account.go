package domain

import (
	"context"
	"cqrs-travian/engine/domain/commands"
	"cqrs-travian/engine/events"
	"fmt"
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

type Account struct {
	*cqrs.AggregateBase
	// current cold
	registered bool
	token      string
	verified   bool
}

func (a *Account) CreateUserAccount(ctx context.Context, cmd *commands.CreateUserAccountCommand) error {
	if a.registered {
		return fmt.Errorf("there already exists an account with this username")
	}
	a.AppendEvent(events.UserAccountCreatedEvent{
		WorldID:  cmd.WorldID,
		Username: cmd.Username,
		Email:    cmd.Email,
		Password: cmd.Password,
	})

	a.AppendEvent(events.UserAccountVerificationTokenCreated{
		WorldID: cmd.WorldID,
		Email:   cmd.Email,
		Token:   uuid.NewString(),
	})

	return nil
}

func (a *Account) OnUserAccountCreated(ev *events.UserAccountCreatedEvent) {
	a.registered = true
}

func (a *Account) UserAccountVerificationToken(ev *events.UserAccountVerificationTokenCreated) {
	a.token = ev.Token
}

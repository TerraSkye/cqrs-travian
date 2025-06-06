package commands

import (
	"fmt"
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var _ cqrs.Command = (*CreateUserAccountCommand)(nil)

type CreateUserAccountCommand struct {
	WorldID  uuid.UUID
	Username string
	Email    string
	Password string
}

func (e *CreateUserAccountCommand) AggregateID() string {
	return fmt.Sprintf("%s-%s", e.WorldID.String(), e.Username)
}

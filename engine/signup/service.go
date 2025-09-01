package signup

import (
	"context"
	"cqrs-travian/engine/domain/commands"
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

type Service interface {
	CreateUserAccount(ctx context.Context, payload Payload) error
}

type Payload struct {
	Username string
	Email    string
	Password string
	WorldID  uuid.UUID
}

type service struct {
	bus cqrs.CommandBus
}

func NewService(bus cqrs.CommandBus) Service {
	return &service{
		bus: bus,
	}
}

func (s service) CreateUserAccount(ctx context.Context, payload Payload) error {
	// price
	cmd := &commands.CreateUserAccountCommand{
		WorldID:  payload.WorldID,
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
	}
	if err := s.bus.Send(ctx, cmd); err != nil {
		return err
	}

	return nil
}

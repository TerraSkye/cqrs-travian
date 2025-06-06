package establishvillage

import (
	"context"
	"cqrs-travian/world/domain/commands"
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

type Service interface {
	EstablishVillage(ctx context.Context, tileID int64) error
}

type Payload struct {
	AggregateId uuid.UUID
	Description string
	Image       string
	ItemId      uuid.UUID
	ProductId   uuid.UUID
}

type service struct {
	bus cqrs.CommandBus
}

func NewService(bus cqrs.CommandBus) Service {
	return &service{
		bus: bus,
	}
}

func (s service) EstablishVillage(ctx context.Context, tileID int64) error {
	// price
	cmd := &commands.EstablishVillageCommand{
		WorldID: uuid.UUID{},
		TileID:  tileID,
	}
	if err := s.bus.Send(ctx, cmd); err != nil {
		return err
	}

	return nil
}

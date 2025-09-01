package domain

import (
	"context"
	"cqrs-travian/engine/domain/commands"
	"cqrs-travian/engine/events"

	cqrs "github.com/terraskye/eventsourcing"
)

type World struct {
	*cqrs.AggregateBase
}

func (a *World) Tick(ctx context.Context, cmd *commands.Tick) error {

	a.AppendEvent(&events.TickEvent{
		WorldID:   cmd.WorldID,
		TickCount: 1,
	})
	return nil
}

func (a *World) OnTick(ev *events.TickEvent) {
	// progressed 1 tick

}

package main

import (
	"context"
	"cqrs-travian/engine"
	"cqrs-travian/engine/domain"
	cmd "cqrs-travian/engine/domain/commands"
	_ "cqrs-travian/engine/handlers"
	world_listing "cqrs-travian/engine/world-listing"
	"cqrs-travian/infra"
	"cqrs-travian/world"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	cqrs "github.com/terraskye/eventsourcing"
)

func main() {

	log := logrus.StandardLogger()

	worldID := uuid.New()

	karte := world.NewWorld(1, 400)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var eventBus cqrs.EventBus

	{
		eventBus = cqrs.NewEventBus()
	}

	var eventStore cqrs.EventStore

	{
		eventStore = cqrs.NewMemoryStore(eventBus)
	}

	var commandHandler cqrs.CommandHandler

	{
		commandHandler = cqrs.NewCommandHandler(
			eventStore,
			engine.AggregateForCommand,
			engine.DispatchEvent,
			engine.DispatchCommand,
		)

		commandHandler = infra.NewLoggingCommandHandler(commandHandler, log.WithField("service", "commandhandler"))
	}

	var commandBus cqrs.CommandBus

	{
		commandBus = cqrs.NewCommandBus(1000)
	}

	commandBus.AddHandler(commandHandler.Handle)

	{
		// world listings

		world_listing.NewQueryHandler(eventStore)

	}

	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Microsecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:

				// push the command, ignoring errors for now
				_ = commandBus.Send(cqrs.WithCausationId(ctx, "ticker"), &cmd.Tick{
					WorldID: worldID,
				})
			}
		}
	}(ctx)

	// Capture OS interrupt signal to cancel context gracefully (optional)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()

	// Block here until context is cancelled
	<-ctx.Done()

	fmt.Println("Context cancelled, exiting")

}

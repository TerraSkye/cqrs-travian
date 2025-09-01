package building_queue

import (
	"context"
	"cqrs-travian/engine/events"
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

var Queue support.Queue[PendingConstruction]

func NewProjector() *cqrs.EventGroupProcessor {

	q := queue{
		Queue: Queue,
	}

	return cqrs.NewEventGroupProcessor("building-queue",
		cqrs.NewGroupEventHandler(q.OnTick),
		cqrs.NewGroupEventHandler(q.OnConstructionStarted),
		cqrs.NewGroupEventHandler(q.OnGoldCompletion),
	)
}

type queue struct {
	Queue support.Queue[PendingConstruction]
}

func (p *queue) OnTick(ctx context.Context, _ *events.TickEvent) error {

	p.Queue.Poll(ctx)
	return nil
}

func (p *queue) OnGoldCompletion(ctx context.Context, ev *events.GoldCompletionEvent) error {

	p.Queue.ExpireAll(ctx, func(p PendingConstruction) bool {
		return ev.VillageID == p.VillageID
	})

	return nil
}

func (p *queue) OnConstructionStarted(ctx context.Context, _ *events.TickEvent) error {

	p.Queue.Add(ctx, 20, PendingConstruction{
		WorldID:        uuid.UUID{},
		VillageID:      uuid.UUID{},
		ConstructionID: uuid.UUID{},
	})
	return nil
}

func (p *queue) OnConstructionCompleted(ctx context.Context, _ *events.TickEvent) error {

	p.Queue.RemoveByID(ctx)

	return nil
}

func (p *queue) OnConstructionCancelled(ctx context.Context, _ *events.TickEvent) error {

	p.Queue.RemoveByID(ctx)

	return nil
}

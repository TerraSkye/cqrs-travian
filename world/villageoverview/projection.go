package villageoverview

import (
	"context"
	"cqrs-travian/world/events"
	"cqrs-travian/world/support"
	cqrs "github.com/terraskye/eventsourcing"
)

var tasks support.Queue[*ReadModel]

func NewProjector() *cqrs.EventGroupProcessor {

	p := &service{}

	return cqrs.NewEventGroupProcessor(
		"village-overview",
		cqrs.NewGroupEventHandler(p.OnTick),
		cqrs.NewGroupEventHandler(p.OnCelebrationStarted),
		cqrs.NewGroupEventHandler(p.OnCelebrationCompleted),
		cqrs.NewGroupEventHandler(p.OnGoldCompletion),
	)
}

type service struct {
}

func (s *service) OnTick(ctx context.Context, ev *events.TickEvent) error {
	tasks.PollN(ctx, ev.TickCount)
	return nil
}

func (s *service) OnCelebrationStarted(ctx context.Context, ev *events.CelebrationStartedEvent) error {
	tasks.Add(ctx, ev.Duration, ev.CelebrationID, &ReadModel{
		AggregateId: ev.CelebrationID,
		VillageID:   ev.VillageID,
		WorldID:     ev.WorldID,
		TTL:         ev.Duration,
		Large:       ev.Large,
	})

	return nil
}

func (s *service) OnCelebrationCompleted(ctx context.Context, ev *events.CelebrationCompletedEvent) error {
	tasks.RemoveByID(ctx, ev.CelebrationID)
	return nil
}

func (s *service) OnGoldCompletion(ctx context.Context, ev *events.CelebrationCompletedEvent) error {
	tasks.ExpireAll(ctx, func(model *ReadModel) bool { return model.VillageID == ev.VillageID })
	return nil
}

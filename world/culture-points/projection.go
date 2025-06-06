package culture_points

import (
	"context"
	"cqrs-travian/world/events"
	cqrs "github.com/terraskye/eventsourcing"
)

func NewProjector() *cqrs.EventGroupProcessor {

	p := &service{}

	return cqrs.NewEventGroupProcessor(
		"culture-points",
		cqrs.NewGroupEventHandler(p.OnTick),
		cqrs.NewGroupEventHandler(p.OnCelebrationStarted),
		cqrs.NewGroupEventHandler(p.OnCelebrationCompleted),
		cqrs.NewGroupEventHandler(p.OnGoldCompletion),
	)
}

type service struct {
	culturePoints map[string]*cqrs.ReadModel
}

func (s *service) OnCelebrationCompleted(ctx context.Context, ev *events.CelebrationCompletedEvent) error {
	s.culturePoints[ev.CulturePoints]
	return nil
}

func (s *service) OnTick(ctx context.Context, ev *events.TickEvent) error {
	s.culturepoints.PollN(ctx, ev.TickCount)
	return nil
}

func (s *service) OnGoldCompletion(ctx context.Context, ev *events.CelebrationCompletedEvent) error {
	tasks.ExpireAll(ctx, func(model *ReadModel) bool { return model.VillageID == ev.VillageID })
	return nil
}

package domain

import (
	"context"
	"cqrs-travian/world/events"
)

type Tutorial struct {
	Account            string
	direction          string // eco / miletary // skipped
	aggregateLifecycle infra.AggregateLifecycle[int64]
	step               int64
}

func (v *Tutorial) OnTaskCompleted(ctx context.Context, event *events.TaskCompletedEvent) {
	v.step++
}

func (v *Tutorial) OnTaskCompleted(ctx context.Context, event *events.TaskCompletedEvent) {
	v.step++
}
